package MgoConn

// add by stefan

import (
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/RedisConn"
	"github.com/Peakchen/xgameCommon/ado"
	. "github.com/Peakchen/xgameCommon/public"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type TAokoMgo struct {
	sync.Mutex

	session     *mgo.Session
	UserName    string
	Passwd      string
	ServiceHost string
	chSessions  chan *mgo.Session
	PoolCnt     int
	server      string
}

func NewMgoConn(server, Username, Passwd, Host string) *TAokoMgo {
	aokomogo := &TAokoMgo{}
	aokomogo.UserName = Username
	aokomogo.Passwd = Passwd
	aokomogo.ServiceHost = Host
	aokomogo.server = server
	//template set 1000 session.
	aokomogo.PoolCnt = int(ado.EMgo_Thread_Cnt)
	aokomogo.chSessions = make(chan *mgo.Session, aokomogo.PoolCnt)
	aokomogo.NewDial()
	return aokomogo
}

func (this *TAokoMgo) NewDial() {
	session, err := this.NewMgoSession()
	if err != nil {
		akLog.FmtPrintln(err)
		return
	}

	this.session = session
	for i := 1; i <= this.PoolCnt; i++ {
		this.chSessions <- session.Copy()
	}
	return
}

func (this *TAokoMgo) NewMgoSession() (session *mgo.Session, err error) {
	MdialInfo := &mgo.DialInfo{
		Addrs:     []string{this.ServiceHost},
		Username:  this.UserName,
		Password:  this.Passwd,
		Direct:    false,
		Timeout:   time.Second * 10,
		PoolLimit: 4096,
		//ReadTimeout: time.Second * 5,
		//WriteTimeout: time.Second * 10,
	}

	session, err = mgo.DialWithInfo(MdialInfo)
	if err != nil {
		err = akLog.RetError("mgo dial err: %v.\n", err)
		akLog.Error("mgo dial err: %v.\n", err)
		return
	}

	err = session.Ping()
	if err != nil {
		err = akLog.RetError("session ping out, err: %v.", err)
		akLog.Error("session ping out, err: %v.", err)
		return
	}

	session.SetMode(mgo.Monotonic, true)
	session.SetCursorTimeout(0)
	// focus on those selects.
	//http://www.mongoing.com/archives/1723
	Safe := &mgo.Safe{
		J:     true,       //true:写入落到磁盘才会返回|false:不等待落到磁盘|此项保证落到磁盘
		W:     1,          //0:不会getLastError|1:主节点成功写入到内存|此项保证正确写入
		WMode: "majority", //"majority":多节点写入|此项保证一致性|如果我们是单节点不需要这只此项
	}
	session.SetSafe(Safe)
	//session.SetSocketTimeout(time.Duration(5 * time.Second()))
	err = nil
	return
}

func (this *TAokoMgo) Exit() {
	if this.chSessions != nil {
		close(this.chSessions)
	}
}

func (this *TAokoMgo) GetMgoSession() (sess *mgo.Session, err error) {
	this.Lock()
	defer this.Unlock()

	if this.session == nil {
		err = fmt.Errorf("aoko mongo session not get invalid.")
		return
	}

	sess = this.session
	err = nil
	return
}

func (this *TAokoMgo) getMgoSessionByChan() (sess *mgo.Session, err error) {
	this.Lock()
	defer this.Unlock()

	select {
	case s, _ := <-this.chSessions:
		return s, nil
	case <-time.After(time.Duration(1) * time.Second):
	default:
	}
	return nil, fmt.Errorf("aoko mongo session time out and not get.")
}

func MakeMgoModel(Identify, MainModel, SubModel string) string {
	return MainModel + "." + SubModel + "." + Identify
}

func (this *TAokoMgo) QueryAcc(usrName string, OutParam IDBCache) (err error, exist bool) {
	condition := bson.M{OutParam.SubModel() + "." + "username": usrName}
	return this.QueryByCondition(condition, OutParam)
}

func (this *TAokoMgo) QueryOne(Identify string, OutParam IDBCache) (err error, exist bool) {
	condition := bson.M{"_id": Identify}
	return this.QueryByCondition(condition, OutParam)
}

func (this *TAokoMgo) QueryByCondition(condition bson.M, OutParam IDBCache) (err error, exist bool) {
	session, err := this.GetMgoSession()
	if err != nil {
		err = akLog.RetError("get sesson err: %v.", err)
		return
	}

	s := session.Clone()
	defer s.Close()

	collection := s.DB(this.server).C(OutParam.MainModel())
	retQuerys := collection.Find(condition)
	count, ret := retQuerys.Count()
	if ret != nil || count == 0 {
		err = akLog.RetError("[mgo] query data err: %v, %v.", ret, count)
		return
	}

	selectRet := retQuerys.Select(bson.M{OutParam.SubModel(): 1, "_id": 1}).Limit(1)
	if selectRet == nil {
		err = akLog.RetError("[mgo] selectRet invalid, submodule: %v.", OutParam.SubModel())
		return
	}

	outval := reflect.MakeMap(reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(OutParam)))
	ret = selectRet.One(outval.Interface())
	if ret != nil {
		err = akLog.RetError("[mgo] select one error: %v.", ret)
		return
	}

	retIdxVal := outval.MapIndex(reflect.ValueOf(OutParam.SubModel()))
	if !retIdxVal.IsValid() {
		err = akLog.RetError("[mgo] outval MapIndex invalid.")
		return
	}

	reflect.ValueOf(OutParam).Elem().Set(retIdxVal.Elem())
	exist = true
	return
}

func (this *TAokoMgo) QuerySome(Identify string, OutParam IDBCache) (err error) {
	session, err := this.GetMgoSession()
	if err != nil {
		return err
	}

	s := session.Clone()
	defer s.Close()

	collection := s.DB(this.server).C(OutParam.MainModel())
	err = collection.Find(bson.M{"_id": Identify}).All(&OutParam)
	if err != nil {
		err = fmt.Errorf("Identify: %v, MainModel: %v, SubModel: %v, err: %v.\n", Identify, OutParam.MainModel(), OutParam.SubModel(), err)
		akLog.Error("[QuerySome] err: %v.\n", err)
	}
	return
}

func (this *TAokoMgo) InsertOne(Identify string, InParam IDBCache) (err error) {
	session, err := this.GetMgoSession()
	if err != nil {
		return err
	}

	s := session.Clone()
	defer s.Close()

	akLog.FmtPrintf("[Insert] main: %v, sub: %v, key: %v.", InParam.MainModel(), InParam.SubModel(), InParam.Identify())
	collection := s.DB(this.server).C(InParam.MainModel())
	operAction := bson.M{"_id": InParam.Identify(), InParam.SubModel(): InParam}
	err = collection.Insert(operAction)
	if err != nil {
		err = fmt.Errorf("main: %v, sub: %v, key: %v, err: %v.\n", InParam.MainModel(), InParam.SubModel(), InParam.Identify(), err)
		akLog.Error("[Insert] err: %v.\n", err)
	}
	return
}

func (this *TAokoMgo) SaveOne(Identify string, InParam IDBCache) (err error) {
	session, err := this.GetMgoSession()
	if err != nil {
		return err
	}

	redkey := MakeMgoModel(InParam.Identify(), InParam.MainModel(), InParam.SubModel())
	err = Save(session, this.server, redkey, InParam)
	return
}

func Save(mgosession *mgo.Session, dbserver, redkey string, data interface{}) (err error) {
	s := mgosession.Clone()
	defer s.Close()

	key, main, sub := RedisConn.ParseRedisKey(redkey)
	akLog.FmtPrintf("update origin: %v, main: %v, sub: %v, key: %v.", redkey, main, sub, key)
	collection := s.DB(dbserver).C(main)
	operAction := bson.M{"$set": bson.M{sub: data}}
	_, err = collection.UpsertId(key, operAction)
	if err != nil {
		err = fmt.Errorf("main: %v, sub: %v, key: %v, err: %v.\n", main, sub, key, err)
		akLog.Error("[Save] err: %v.\n", err)
	}
	return
}

func (this *TAokoMgo) EnsureIndex(InParam IDBCache, idxs []string) (err error) {
	session, err := this.GetMgoSession()
	if err != nil {
		return err
	}
	s := session.Clone()
	defer s.Close()

	c := s.DB(this.server).C(InParam.SubModel())
	err = c.EnsureIndex(mgo.Index{Key: idxs})
	return err
}
