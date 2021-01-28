package mysql

import (
	"database/sql"

	"github.com/Peakchen/xgameCommon/akLog"
)

func LoadMysqlDB(dbpath, dbName string) (mydb *AkMysql) {
	db, err := sql.Open("mysql", dbpath+dbName)
	if err != nil {
		akLog.Error("load mysql db fail, err: ", err)
		return
	}

	if err := db.Ping(); err != nil {
		akLog.Error("mysql db ping fail, err: ", err)
		return
	}

	return NewAkMysql(db)
}
