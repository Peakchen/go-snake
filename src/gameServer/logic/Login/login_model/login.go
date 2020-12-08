package login_model

import (
	"github.com/Peakchen/xgameCommon/akLog"
	"gorm.io/gorm"
)

type User struct {
	UID string `xorm:"UID" pk`
}

func (this *User) LoadUsers(session *gorm.Session) []*User {
	var users []*User
	err := session.Find(this, &users)
	if err != nil {
		akLog.Fail(err)
		return nil
	}
	return users
}
