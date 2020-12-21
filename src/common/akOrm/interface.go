package akOrm

type IAkModel interface {
	GetUserID() int64
	Update() bool
	Delete() bool
	Create() bool
}
