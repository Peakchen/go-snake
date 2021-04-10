package svrbalance

// server balance define

type TExternal struct {
	Persons int32
}

type IBalance interface {
	AddSvr(svr string)
	Push(svr string)
	GetSvr() (s string)
}

const (
	ESvrBalanceMaxPersons = int32(10000)
)
