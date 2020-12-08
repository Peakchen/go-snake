package dbEngine

type OpType uint8

const (
	OP_INSERT = OpType(1)
	OP_UPDATE = OpType(2)
	OP_DELETE = OpType(3)
	OP_QUERY  = OpType(4)
)

const (
	maxDBOpQueueSize = 1000
)

const (
	DRIVER_MYSQL   = "mysql"
	DRIVER_SQLITE3 = "sqlite3"
)
