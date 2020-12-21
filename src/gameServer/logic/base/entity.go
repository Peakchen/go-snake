package base

type IEntity interface {
	GetID() int64
	SetID(id int64)
	GetSessionID() string
	SetSessionID(id string)
}

type IEntityUser interface {
	//self

	//child
	IEntity
	IDevice
}

type IEntityAI interface {
	//self

	//child
	IEntity
}
