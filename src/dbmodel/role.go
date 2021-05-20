package dbModule

import (
	"go-snake/common/akOrm"
	"reflect"

	"github.com/Peakchen/xgameCommon/utils"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model

	RowID  string `gorm:"column:rowid;type:varchar(50);primary_key" bson:"rowid" json:"rowid"`     //数据操作行ID
	DBID   int64  `gorm:"column:dbid;unique_index" bson:"dbid" json:"dbid"`                        //数据库唯一ID
	UserID string `gorm:"column:userid;type:varchar(50);unique_index" bson:"userid" json:"userid"` //玩家ID 供客户端显示

	Level    uint32 `gorm:"colum:level;type:uint" `
	Name     string `gorm:"colum:name;type:string;default:''" `
	HeadIcon string `gorm:"colum:headicon;type:string;default:''" `
}

func (this *Role) TableName() string {
	return reflect.TypeOf(*this).Name()
}

func NewRole(uid string, dbid int64, level uint32, name string, headicon string) *Role {
	role := &Role{
		RowID:    utils.GetOnlyString_v6(),
		DBID:     dbid,
		UserID:   uid,
		Level:    level,
		Name:     name,
		HeadIcon: headicon,
	}
	if !akOrm.Create(role) {
		return nil
	}
	return role
}

func LoadRoles() []*Role {
	var rets []*Role
	akOrm.Find(&rets)
	return rets
}

func LoadRole(id int64)*Role{
	var ret Role
	akOrm.FindOne(&ret)
	return &ret
}

func (this *Role) GetDBID() int64 {
	return this.DBID
}

func (this *Role) Create() bool {
	return akOrm.Create(this)
}

func (this *Role) Delete() bool {
	return akOrm.Delete(this)
}

func (this *Role) Update() bool {
	return akOrm.Update(this)
}
