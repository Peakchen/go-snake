package accdb

import (
	"go-snake/common/akOrm"
	"reflect"

	"github.com/Peakchen/xgameCommon/utils"

	"gorm.io/gorm"
)

type Acc struct {
	gorm.Model

	RowID  string `gorm:"column:rowid;type:varchar(50);primary_key" bson:"rowid" json:"rowid"`     //数据操作行ID
	DBID   int64  `gorm:"column:dbid;unique_index" bson:"dbid" json:"dbid"`                        //数据库唯一ID
	UserID string `gorm:"column:userid;type:varchar(50);unique_index" bson:"userid" json:"userid"` //玩家ID 供客户端显示

	User string `gorm:"column:user;type:varchar(10);unique_index" bson:"user" json:"user"`
	Pwd  string `gorm:"column:pwd;type:varchar(20);unique_index" bson:"pwd" json:"pwd"`
}

func (this *Acc) TableName() string {
	return reflect.TypeOf(*this).Name()
}

func (this *Acc) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (this *Acc) BeforeUpdate(tx *gorm.DB) (err error) {
	return nil
}

func (this *Acc) BeforeSave(tx *gorm.DB) (err error) {
	return nil
}

func (this *Acc) AfterFind(tx *gorm.DB) (err error) {
	return nil
}

func (this *Acc) AfterCreate(tx *gorm.DB) (err error) {
	return nil
}

func (this *Acc) AfterUpdate(tx *gorm.DB) (err error) {
	return nil
}

func (this *Acc) AfterSave(tx *gorm.DB) (err error) {
	return nil
}

func (this *Acc) BeforeDelete(tx *gorm.DB) (err error) {
	return nil
}

func (this *Acc) AfterDelete(tx *gorm.DB) (err error) {
	return nil
}

func NewAcc(userName string, pwd string) *Acc {
	acc := &Acc{
		RowID:  utils.GetOnlyString_v6(),
		DBID:   utils.NewInt64_v1(),
		UserID: utils.GetOnlyString_v4(),
		User:   userName,
		Pwd:    pwd,
	}
	if !akOrm.Create(acc) {
		return nil
	}
	return acc
}

func (this *Acc) Load() []*Acc {
	var rets []*Acc
	akOrm.Find(&rets)
	return rets
}

func (this *Acc) GetDBID() int64 {
	return this.DBID
}

func (this *Acc) Create() bool {
	return akOrm.Create(this)
}

func (this *Acc) Delete() bool {
	return akOrm.Delete(this)
}

func (this *Acc) Update() bool {
	return akOrm.Update(this)
}
