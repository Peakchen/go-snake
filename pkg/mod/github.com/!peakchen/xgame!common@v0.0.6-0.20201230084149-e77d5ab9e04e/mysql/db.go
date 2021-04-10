package mysql

import (
	"common"
	"database/sql"
	"fmt"
)

type TQueryField struct {
	Field string
	Data  interface{}
}

type AkMysql struct {
	mysqlDB *sql.DB
}

func NewAkMysql(db *sql.DB) *AkMysql {
	return &AkMysql{
		mysqlDB: db,
	}
}

func (this *AkMysql) Update(sqlcmds string) (err error) {
	_, err = this.mysqlDB.Exec(sqlcmds)
	if err != nil {
		err = fmt.Errorf("db Exec update, err: %v.", err)
	}
	return
}

func (this *AkMysql) Query(sql string, ret *[]([]*TQueryField)) (err error) {
	rows, err := this.mysqlDB.Query(sql)
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
