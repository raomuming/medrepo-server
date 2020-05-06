package model

import (
	"errors"
	"log"

	//"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"

	"medrepo-server/middleware"
	"medrepo-server/util/aes"
)

type User struct {
	Model
	Verify int
	Wechat Wechat
}

// Login
func (u *User) Login() (string, error) {
	if u.Wechat.Openid == "" || u.Wechat.SessionKey == "" {
		return "", errors.New("用户信息不完整")
	}

	DB().Where("openid=?", u.Wechat.Openid).Find(&u.Wechat)
	uid := u.Wechat.UserID
	if uid == 0 {
		if err := DB().Create(u).Error; err != nil {
			return "", err
		}
		uid = u.ID
	}

	DB().Find(u, uid)

	return middleware.CreateToken(uid)
}

func (u *User) AfterCreate(scope *gorm.Scope) error {
	DB().Create(&UserConfig{
		UserID: u.ID,
		Notify: NotifyAll,
	})
	return nil
}

func (u *User) BeforeSave(scope *gorm.Scope) error {
	return nil
}