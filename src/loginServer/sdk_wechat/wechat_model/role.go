package wechat_model

import (
	"database/sql/driver"
	"go-snake/common/akOrm"
	"go-snake/common/wechat"
	"reflect"

	"github.com/Peakchen/xgameCommon/utils"

	"gorm.io/gorm"
)

type TWaterMark struct {
	Timestamp int64  `gorm:"column:timestamp" json:"timestamp"`
	AppID     string `gorm:"column:appid;type:varchar(50)" json:"appid"`
}

func (wm *TWaterMark) Scan(value interface{}) (err error) {
	wm = value.(*TWaterMark)
	return
}

func (wm TWaterMark) Value() (driver.Value, error) {
	return wm, nil
}

func (wm TWaterMark) GormDataType() string {
	return "wm"
}

type WxRole struct {
	gorm.Model

	DBID      int64       `gorm:"column:dbid;primary_key" json:"dbid"`          //数据库id primarykey
	OpenID    string      `gorm:"column:OpenID;type:varchar(50)" json:"openId"` //用户唯一标识符
	UnionID   string      `gorm:"column:union;type:varchar(50)" json:"unionId"` //
	NickName  string      `gorm:"column:nickname;type:varchar(100)" json:"nickName"`
	Gender    int         `gorm:"column:gender;type:int" json:"gender"`
	City      string      `gorm:"column:city;type:varchar(20)" json:"city"`
	Province  string      `gorm:"column:province;type:varchar(20)" json:"province"`
	Country   string      `gorm:"column:country;type:varchar(20)" json:"country"`
	AvatarURL string      `gorm:"column:avatarurl;type:varchar(200)" json:"avatarUrl"`
	Language  string      `gorm:"column:language;type:varchar(50)" json:"language"`
	Watermark *TWaterMark `gorm:"column:watermark;embedded" json:"watermark"`
}

func (this *WxRole) TableName() string {
	return reflect.TypeOf(*this).Name()
}

func (this *WxRole) BeforeCreate(tx *gorm.DB) {
	field := tx.Statement.Schema.LookUpField("watermark")
	if field.DataType == "wm" {
		this.Watermark = &TWaterMark{}
	}
}

func (this *WxRole) Copy(userinfo *wechat.WxUserInfo) {
	this.OpenID = userinfo.OpenID
	this.UnionID = userinfo.UnionID
	this.NickName = userinfo.NickName
	this.Gender = userinfo.Gender
	this.City = userinfo.City
	this.Province = userinfo.Province
	this.Country = userinfo.Country
	this.AvatarURL = userinfo.AvatarURL
	this.Language = userinfo.Language
	this.Watermark.Timestamp = userinfo.Watermark.Timestamp
	this.Watermark.AppID = userinfo.Watermark.AppID
}

func NewRoleAtWX(userinfo *wechat.WxUserInfo) *WxRole {
	acc := &WxRole{
		DBID:      utils.NewInt64_v1(),
		Watermark: &TWaterMark{},
	}
	acc.Copy(userinfo)
	if !akOrm.Create(acc) {
		return nil
	}
	return acc
}

func (this *WxRole) Load() []*WxRole {
	var rets []*WxRole
	akOrm.Find(&rets)
	return rets
}

func (this *WxRole) GetDBID() int64 {
	return this.DBID
}

func (this *WxRole) Create() bool {
	return akOrm.Create(this)
}

func (this *WxRole) Delete() bool {
	return akOrm.Delete(this)
}

func (this *WxRole) Update() bool {
	return akOrm.Update(this)
}
