package model

import (
	"errors"

	"medrepo-server/middleware"
	"medrepo-server/mlog"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Wechat Wechat
}

func (u *User) Login() (string, error) {
	if u.Wechat.Openid == "" || u.Wechat.SessionKey == "" {
		mlog.Error("user info is not complete")
		return "", errors.New("user info is not complete")
	}

	DB().Where("openid=?", u.Wechat.Openid).Find(&u.Wechat)
	uid := u.Wechat.UserID
	if uid == 0 {
		if err := DB().Create(u).Error; err != nil {
			mlog.Error("create user failed")
			return "", err
		}
		uid = u.ID
	}

	DB().Find(u, uid)
	return middleware.CreateToken(uid)
}

func (u *User) AfterCreate(scope *gorm.Scope) error {
	return nil
}

func (u *User) BeforeSave(scope *gorm.Scope) error {
	return nil
}

func (u *User) AfterFind(scope *gorm.Scope) error {
	return nil
}

func GetCollector(userID uint) (*colly.Collector, error) {
	user := User{}
	if err := DB().Find(&user, userID).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func RemoveAll(uid uint) {
	del := DB().Unscoped().Where("user_id = ?", uid).Delete
	del(Wechat{})
	DB().Unscoped().Where("id = ?", uid).Delete(User{})
}
