package akOrm

import (
	"fmt"
	"go-snake/common"
	"runtime"
	"sync"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBEngine struct {
	sync.RWMutex

	connCfg string
	ormDB   *gorm.DB
	actor   *DBActor
}

var (
	_dbengines = make(map[int64]*DBEngine)
	maxlp      = int64(runtime.NumCPU())
	_db        *gorm.DB
	dbCfg      string
)

func newDB(cfg string) (db *gorm.DB) {
	var err error
	db, err = gorm.Open(mysql.Open(cfg), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		akLog.Error("open mysql fail,config: ", cfg)
	}
	return db
}

func OpenDB(user, pwd, host, dbName string) {
	common.DosafeRoutine(func() {
		dbCfg = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, host, dbName)
		_db = newDB(dbCfg)
		if _db == nil {
			panic("exit.")
		}
		/*
			无服务监听时，注册则会报错
			db.Use(prometheus.New(prometheus.Config{
				DBName:          dbName,                      // 使用 `DBName` 作为指标 label
				RefreshInterval: 15,                          // 指标刷新频率（默认为 15 秒）
				PushAddr:        "prometheus pusher address", // 如果配置了 `PushAddr`，则推送指标
				StartServer:     true,                        // 启用一个 http 服务来暴露指标
				HTTPServerPort:  8080,                        // 配置 http 服务监听端口，默认端口为 8080 （如果您配置了多个，只有第一个 `HTTPServerPort` 会被使用）
				MetricsCollector: []prometheus.MetricsCollector{
					&prometheus.MySQL{
						VariableNames: []string{"Threads_running"},
					},
				}, // 用户自定义指标
			}))
		*/
		for i := int64(0); i < maxlp; i++ {
			_dbengines[i] = &DBEngine{
				connCfg: dbCfg,
				ormDB:   _db,
			}
			sdb, err := _dbengines[i].ormDB.DB()
			if err != nil || sdb == nil {
				fmt.Println(err)
			}
			_dbengines[i].actor = newDBActor(_dbengines[i])
			_dbengines[i].actor.loop()
		}
	}, func() {
		time.Sleep(time.Second)
		common.SafeExit()
	})
}

func (this *DBEngine) checkConnect() bool {
	reconnect := func() bool {
		this.Lock()
		defer this.Unlock()

		this.ormDB = nil
		var reconns int
		for this.ormDB == nil && reconns < 3 {
			db := newDB(this.connCfg)
			if db != nil {
				break
			}
			this.ormDB = db
			time.Sleep(time.Duration(1) * time.Second)
			reconns++
		}
		if this.ormDB == nil {
			return false
			//common.SafeExit()
		}
		return true
	}
	db, err := this.ormDB.DB()
	if err != nil {
		return reconnect()
	} else if err := db.Ping(); err != nil {
		return reconnect()
	}
	return true
}

func GetDBActor(rowID int64) *DBActor {
	en, ok := _dbengines[rowID%maxlp]
	if !ok {
		return nil
	}
	return en.actor
}

type DBOper struct {
	Type int
	Data IAkModel
}

type DBActor struct {
	*DBEngine
	oper chan *DBOper
	stop chan bool
	wg   sync.WaitGroup
}

func newDBActor(db *DBEngine) *DBActor {
	return &DBActor{
		DBEngine: db,
		oper:     make(chan *DBOper, 1000),
		stop:     make(chan bool),
	}
}

func (this *DBActor) loop() {

	common.DosafeRoutine(func() {
	dbloop:
		for {
			select {
			case oper := <-this.oper:
				var sess *gorm.DB
				common.Dosafe(func() {
					if !this.checkConnect() {
						this.stop <- true
					} else {
						sess = this.ormDB.Session(&gorm.Session{PrepareStmt: true})
					}
				}, nil)
				if sess == nil {
					break dbloop
				}
				switch oper.Type {
				case ORM_CREATE:
					common.Dosafe(func() {
						sess.Create(oper.Data)
						this.oper <- &DBOper{
							Type: ORM_UPDATE,
							Data: oper.Data,
						}
					}, func() {
						this.stop <- true
					})
				case ORM_UPDATE:
					common.Dosafe(func() { sess.Save(oper.Data) }, func() {
						this.stop <- true
					})
				case ORM_DELETE:
					common.Dosafe(func() { sess.Delete(oper.Data) }, func() {
						this.stop <- true
					})
				}
			case <-this.stop:
				break dbloop
			}
		}
		this.wg.Done()
	}, func() {
		time.Sleep(time.Second)
	})

}

func (this *DBActor) Do(oper int, m IAkModel) {
	this.oper <- &DBOper{
		Type: oper,
		Data: m,
	}
}

func (this *DBActor) DB() *gorm.DB {
	if !this.checkConnect() {
		return nil
	}
	return this.ormDB
}
