package mysql

import (
	"database/sql"

	"github.com/Peakchen/xgameCommon/akLog"
)

func LoadSqliteDB(dbpath, dbName string) (sqlite *AkSqlite) {
	db, err := sql.Open("sqlite3", dbpath+dbName)
	if err != nil {
		akLog.Error("load sqlite db fail, err: ", err)
		return
	}

	if err := db.Ping(); err != nil {
		akLog.Error("sqlite db ping fail, err: ", err)
		return
	}

	return NewAkMysql(db)
}

//"github.com/jinzhu/gorm"
func LoadgormDB(dbpath, dbName string) {

}
