package ado

// add by stefan

// some data could be saved to cache or db for Persistence.
type EDataUpdateType int32

const (
	ED_SYNC_DB    EDataUpdateType = 1
	ED_SYNC_REDIS EDataUpdateType = 2
)

const (
	EDB_DATA_SAVE_INTERVAL int32 = 10 //5 * 60
)

const (
	EMgo_Thread_Cnt = int32(1000)
)

/*

 */
type IDBModule struct {
	StrIdentify string `bson:"_id" json:"_id"`
}

type EDBOperType int32

const (
	EDBOper_Insert EDBOperType = 1
	EDBOper_Update EDBOperType = 2
	EDBOper_Delete EDBOperType = 3
	EDBOper_DB     EDBOperType = 4
	EDBOper_EXPIRE EDBOperType = 5
	EDBOper_LAND   EDBOperType = 6
	// ...
)

const (
	EDBMgoOper_Update = string("MgoUpdate")
)

type TRedisConfig struct {
	DBIndex  int32
	Connaddr string
	Passwd   string

	Shareconnaddr string
	Sharedbindex  int32
	Pprofaddr     string
}

type TMgoConfig struct {
	Username string
	Passwd   string
	Host     string

	Shareusername string
	Sharepasswd   string
	Sharehost     string

	Pprofaddr string
}
