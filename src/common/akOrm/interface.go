package akOrm

type IAkModel interface {
	GetDBID() int64
	Update() bool
	Delete() bool
	Create() bool
}
