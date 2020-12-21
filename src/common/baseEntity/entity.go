package baseEntity

type IEntity interface {
	GetID() int64
	SetID(id int64)
	GetSessionID() string
	SetSessionID(id string)
}
