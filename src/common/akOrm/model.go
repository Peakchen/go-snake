package akOrm

import (
	"reflect"

	"github.com/Peakchen/xgameCommon/akLog"
)

//尝试兼容mongo和mysql存储
type AkModel struct {
	RowID     string `gorm:"primary_key" bson:"_id" json:"_id"`
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
}

func Update(m IAkModel) bool {
	if !hasExist(m) {
		return false
	}
	actor := GetDBActor(m.GetUserID())
	if actor == nil {
		akLog.Error("can not find db actor, rowid: ", m.GetUserID())
		return false
	}
	actor.Do(ORM_UPDATE, m)
	return true
}

func Delete(m IAkModel) bool {
	if !hasExist(m) {
		return false
	}
	actor := GetDBActor(m.GetUserID())
	if actor == nil {
		akLog.Error("can not find db actor, userid: ", m.GetUserID())
		return false
	}
	actor.Do(ORM_DELETE, m)
	return true
}

func Create(m IAkModel) bool {
	if hasExist(m) {
		return false
	}
	actor := GetDBActor(m.GetUserID())
	if actor == nil {
		akLog.Error("can not find db actor, userid: ", m.GetUserID())
		return false
	}
	actor.Do(ORM_CREATE, m)
	return true
}

func checkRec() bool {
	rc := func() bool {
		_db = newDB(dbCfg)
		return _db == nil
	}
	db, err := _db.DB()
	if err != nil {
		if rc() {
			akLog.Fail("invalid db.")
			return false
		}
	} else if err := db.Ping(); err != nil {
		if rc() {
			akLog.Fail("can not connect db.")
			return false
		}
	}
	return true
}

func hasExist(m IAkModel) bool {
	if !checkRec() {
		return false
	}
	ref := reflect.New(reflect.TypeOf(m).Elem()).Interface()
	_db.AutoMigrate(ref)
	_db.First(ref, "userid = ?", m.GetUserID())
	if ref == nil {
		return false
	}
	im := ref.(IAkModel)
	return im.GetUserID() != 0
}

func HasExistAcc(m IAkModel, user string, pwd string) bool {
	if !checkRec() {
		return false
	}
	_db.AutoMigrate(m)
	_db.First(m, "user = ?", user, "pwd = ?", pwd)
	return m.GetUserID() != 0
}

func Find(ms interface{}) {
	if !checkRec() {
		return
	}
	_db.Find(ms)
}
