package base

type IEntity interface {
	GetID() string
	SetID(id string)
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
