package mysql

import (
	"common"
	"database/sql"
	"fmt"

	"github.com/Peakchen/xgameCommon/akLog"
)

type TQueryField struct {
	Field string
	Data  interface{}
}

type AkSqlite struct {
	sqliteDB *sql.DB
}

func NewAkMysql(db *sql.DB) *AkSqlite {
	return &AkSqlite{
		sqliteDB: db,
	}
}

func (this *AkSqlite) Close() {
	this.sqliteDB.Close()
}

func (this *AkSqlite) Update(sql string) (err error) {
	_, err = this.sqliteDB.Exec(sql)
	if err != nil {
		err = fmt.Errorf("db Exec a sql, err: %v.", err)
	}
	return
}

func (this *AkSqlite) Query(sql string, ret *[]([]*TQueryField)) (err error) {
	rows, err := this.sqliteDB.Query(sql)
	if err != nil {
		err = fmt.Errorf("db prepare query, err: %v.", err)
		return
	}

	cols, err := rows.Columns()
	if err != nil {
		err = fmt.Errorf("Columns: %v", err)
		return
	}

	fieldsdata := make([]interface{}, len(cols))
	for i := range fieldsdata {
		var temp interface{}
		fieldsdata[i] = &temp
	}

	for rows.Next() {
		err = rows.Scan(fieldsdata...)
		if err != nil {
			err = fmt.Errorf("db scan field data fail, err: %v.", err)
			return
		}

		columdata := []*common.TQueryField{}
		for idx, data := range fieldsdata {
			akLog.FmtPrintf("Field: ", cols[idx], ", data: ", *data.(*interface{}))
			columdata = append(columdata, &common.TQueryField{
				Field: cols[idx],
				Data:  *data.(*interface{}),
			})
		}

		*ret = append(*ret, columdata)
	}

	if err = rows.Close(); err != nil {
		err = fmt.Errorf("error closing rows: %s", err)
		return
	}

	err = nil
	return
}
