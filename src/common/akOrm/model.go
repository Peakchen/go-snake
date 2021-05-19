package akOrm

import (
	"errors"
	"reflect"

	"github.com/Peakchen/xgameCommon/akLog"
)

func Update(m IAkModel) bool {
	if !hasExist(m) {
		return false
	}
	actor := GetDBActor(m.GetDBID())
	if actor == nil {
		akLog.Error("can not find db actor, rowid: ", m.GetDBID())
		return false
	}
	actor.Do(ORM_UPDATE, m)
	return true
}

func Delete(m IAkModel) bool {
	if !hasExist(m) {
		return false
	}
	actor := GetDBActor(m.GetDBID())
	if actor == nil {
		akLog.Error("can not find db actor, userid: ", m.GetDBID())
		return false
	}
	actor.Do(ORM_DELETE, m)
	return true
}

func Create(m IAkModel) bool {
	if hasExist(m) {
		return false
	}
	actor := GetDBActor(m.GetDBID())
	if actor == nil {
		akLog.Error("can not find db actor, userid: ", m.GetDBID())
		return false
	}
	actor.Do(ORM_CREATE, m)
	return true
}

func checkDB() bool {
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
	if !checkDB() {
		return false
	}
	ref := reflect.New(reflect.TypeOf(m).Elem()).Interface()
	_db.AutoMigrate(ref)
	_db.First(ref, "userid = ?", m.GetDBID())
	if ref == nil {
		return false
	}
	im := ref.(IAkModel)
	return im.GetDBID() != 0
}

func HasExistAcc(m IAkModel, user string, pwd string) (bool, error) {
	if !checkDB() {
		return false, errors.New("db session disconnect.")
	}
	_db.AutoMigrate(m)
	_db.First(m, "user = ?", user, "pwd = ?", pwd)
	return m.GetDBID() != 0, nil
}

func HasExistAcc2(m IAkModel, acc string) (bool, error) {
	if !checkDB() {
		return false, errors.New("db session disconnect.")
	}
	_db.AutoMigrate(m)
	_db.First(m, "account = ?", acc)
	return m.GetDBID() != 0, nil
}

func HasExistForWx(m IAkModel, openid string) (bool, error) {
	
	if !checkDB() {
		return false, errors.New("db session disconnect.")
	}
	
	_db.AutoMigrate(m)
	_db.First(m, "openid = ?", openid)
	
	return m.GetDBID() != 0, nil
	
}

func GetBackUser(m IAkModel, acc, pwd string) (bool, error) {

	if !checkDB() {
		return false, errors.New("db session disconnect.")
	}

	_db.AutoMigrate(m)
	_db.First(m, "account = ?", acc, "password = ?", pwd)
	
	return m.GetDBID() != 0, nil

}

func GetModel(m interface{}, openid string) error {
	if !checkDB() {
		return errors.New("invalid db status.")
	}
	_db.AutoMigrate(m)
	_db.First(m, "openid = ?", openid)
	return nil
}

func Find(ms interface{}) {
	if !checkDB() {
		return
	}
	_db.Find(ms)
}

func FindOne(m IAkModel){

	if !checkDB() {
		return
	}

	ref := reflect.New(reflect.TypeOf(m).Elem()).Interface()
	_db.AutoMigrate(ref)
	_db.First(ref, "userid = ?", m.GetDBID())
	
}