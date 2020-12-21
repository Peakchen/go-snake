package dbEngine

import (
	"go-snake/common"
	"sync"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

type IDataBase interface {
	Update()
	Delete()
}

type DataBases []IDataBase

type Op struct {
	op   OpType
	data []interface{}
}

type DBEngine struct {
	sync.Mutex

	OpCh    chan *Op
	db      *xorm.Engine
	exitCh  chan bool
	session *xorm.Session
}

var _dbEngine *DBEngine

func startDBEngine(dbSrc, dbName string) {
	common.Dosafe(func() {
		_dbEngine = NewDBEngine()
		var err error
		_dbEngine.db, err = xorm.NewEngine(dbSrc, dbName)
		if err != nil {
			akLog.Error("create db fail from file.", dbSrc, dbName)
			return
		}
		tzl, _ := time.LoadLocation("Asia/Shanghai")
		_dbEngine.db.SetTZLocation(tzl)
		_dbEngine.db.SetTZDatabase(tzl)
		_dbEngine.db.ShowSQL(true)
		_dbEngine.db.SetMaxIdleConns(3)
		_dbEngine.db.SetMaxOpenConns(3)

		tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "ak_")
		_dbEngine.db.SetTableMapper(tbMapper)
		_dbEngine.db.ShowExecTime(true)
		_dbEngine.db.Logger().SetLevel(core.LOG_ERR)
		_dbEngine.session = _dbEngine.db.NewSession()
		if err := _dbEngine.session.Ping(); err != nil {
			akLog.FmtPrintln("ping err: ", err)
			return
		}

		go _dbEngine.loop()
		//<-_dbEngine.exitCh
	}, nil)
}

func NewDBEngine() *DBEngine {
	engine := &DBEngine{}
	engine.Init()
	return engine
}

func (this *DBEngine) Init() {
	this.OpCh = make(chan *Op, maxDBOpQueueSize)
	this.exitCh = make(chan bool, 1)
}

func (this *DBEngine) TableExist(obj interface{}) (bool, error) {
	return this.db.IsTableExist(obj)
}

func (this *DBEngine) NewTable(tables []interface{}) {
	this.Lock()

	defer func() {
		this.Unlock()
	}()

	err := this.db.Sync2(tables...)
	if err != nil {
		akLog.FmtPrintln("create table fail, err: ", err)
	}
}

func (this *DBEngine) Insert(obj []interface{}) {
	this.OpCh <- &Op{
		op:   OP_INSERT,
		data: obj,
	}
}

func (this *DBEngine) Update(obj []interface{}) {
	this.OpCh <- &Op{
		op:   OP_UPDATE,
		data: obj,
	}
}

func (this *DBEngine) Delete(obj []interface{}) {
	this.OpCh <- &Op{
		op:   OP_DELETE,
		data: obj,
	}
}

func (this *DBEngine) Find(obj interface{}) (err error) {
	err = this.db.Find(obj)
	return
}

func (this *DBEngine) loop() {
	defer func() {
		this.exitCh <- true
	}()

	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			//do something...
			akLog.FmtPrintln("ticker,t: ", time.Now().Unix())
		case d := <-this.OpCh:
			akLog.FmtPrintln("op data,t: ", d.op)
			switch d.op {
			case OP_INSERT:
				_, err := this.db.Insert(d.data...)
				if err != nil {
					akLog.Error("db Insert fail,err: ", err)
				}
			case OP_UPDATE:
				_, err := this.db.Update(d.data[0])
				if err != nil {
					akLog.Error("db update fail,err: ", err)
				}
			case OP_DELETE:
				_, err := this.db.Delete(d.data[0])
				if err != nil {
					akLog.Error("db delete fail,err: ", err)
				}
			case OP_QUERY:
				akLog.FmtPrintln("there has query db???")
			default:
				akLog.Error("error db op: %v.", d.op)
			}
		}
	}
}

func TableExisted(obj interface{}) bool {
	if ok, err := _dbEngine.TableExist(obj); !ok {
		akLog.Error("err: ", err)
		return false
	}
	return true
}

func Create(objs []interface{}) {
	_dbEngine.NewTable(objs)
}

func Insert(obj []interface{}) {
	_dbEngine.Insert(obj)
}

func Update(obj []interface{}) {
	_dbEngine.Update(obj)
}

func Delete(obj []interface{}) {
	_dbEngine.Delete(obj)
}

func Load(objs interface{}) (err error) {
	err = _dbEngine.Find(objs)
	return
}
