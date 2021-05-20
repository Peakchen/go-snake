package backdb

import (
	"reflect"
	"go-snake/common/akOrm"
	"gorm.io/gorm"
)

/*
	后台用户管理
*/

type User struct {
	gorm.Model

	DBID int64  		`gorm:"column:dbid;unique_index" bson:"dbid" json:"dbid"`
	Account string		`gorm:"column:account;unique_index" bson:"account" json:"account"`
	Password string		`gorm:"column:password" bson:"password" json:"password"`
}

func (this *User) TableName() string {
	return reflect.TypeOf(*this).Name()
}

func (this *User) GetDBID() int64 {
	return this.DBID
}

func (this *User) Create() bool {
	return akOrm.Create(this)
}

func (this *User) Delete() bool {
	return akOrm.Delete(this)
}

func (this *User) Update() bool {
	return akOrm.Update(this)
}

func NewUser(id int64, acc, pwd string)*User{
	return &User{
		DBID: id,
		Account: acc,
		Password: pwd,
	}
}

func LoadUserByID(id int64)*User{
	var ret User
	akOrm.FindOne(&ret)
	return &ret
}

func LoadUser(acc, pwd string)*User{
	var ret User
	_, _ = akOrm.GetBackUser(&ret, acc, pwd)
	return &ret
}

func FindUser(acc string)bool {
	var ret User
	exist,_ := akOrm.HasExistAcc2(&ret, acc)
	return exist
}
