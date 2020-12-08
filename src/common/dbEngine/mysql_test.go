package dbEngine

import (
	"testing"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"
)

func init() {
	akLog.FmtPrintln("init...")
	StartMysql("root:OPbySYJ1FHfl9376ybgzZHtWqt2rcA51@/test?charset=utf8") //tcp(localhost:3306)
}

type User struct {
	IDataBase `json: "-" xorm: "-"`

	UserId  int64     `json:"userId" xorm:"BIGINT(11) pk 'userId' autoincr"`
	Name    string    `json:"name" xorm: "varchar(50) notnull unique 'name'"`
	Addtime time.Time `json:"addtime" xorm: "created"`
	Version int       `json:"version" xorm: "INT(16)"`
}

func (this *User) TableName() string {
	return "ak_user"
}

func LoadTest() {
	var data []*User
	err := Load(&data)
	if err != nil {
		akLog.FmtPrintln("load data fail, err: ", err)
		return
	}
	for _, item := range data {
		akLog.FmtPrintln("load item: ", item.UserId, item.Name, item.Addtime, item.Version)
	}
}

func (this *User) Update() {
	Update([]interface{}{this})
}

func (this *User) Delete() {
	Delete([]interface{}{this})
}

type TA struct {
	IDataBase `xorm: "-"`

	Id      int64     `xorm: "not null pk  autoincr BIGINT(11)"`
	Addtime time.Time `xorm: "DATETIME"`
	Version int       `xorm:"version"`
}

func (this *TA) Update() {
	Update([]interface{}{this})
}

func (this *TA) Delete() {
	Delete([]interface{}{this})
}

func TestMysql(t *testing.T) {
	akLog.FmtPrintln("mysql test")

	Create([]interface{}{new(User)})
	Insert([]interface{}{&User{
		Name:    "1",
		Addtime: time.Now(),
	}})
	time.Sleep(time.Second)
	Insert([]interface{}{&User{
		Name:    "2",
		Addtime: time.Now(),
	}})
	LoadTest()
	var ch = make(chan bool)
	<-ch

	akLog.FmtPrintln("mysql test end.")
}

func TestSelectTable(t *testing.T) {
	if !TableExisted(new(User)) {
		Create([]interface{}{new(User)})
	}
	Insert([]interface{}{
		&User{
			Name:    "3",
			Addtime: time.Now(),
		},
		&User{
			Name:    "4",
			Addtime: time.Now(),
		},
		&User{
			Name:    "5",
			Addtime: time.Now(),
		},
	})

	time.Sleep(time.Second)
	LoadTest()
}
