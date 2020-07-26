package model

import (
	"fmt"
	"time"

	"medrepo-server/config"
	"medrepo-server/mlog"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func initDB() {
	conf := config.Get().Mysql

	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.DB))
	if err != nil {
		mlog.Error("connect to db error")
	}
	db.DB().SetMaxOpenConns(20)
	db.DB().SetConnMaxLifetime(10 * time.Second)
	db.LogMode(config.Get().Debug)

	autoMigrate()
}

func autoMigrate() {
	db.AutoMigrate(
		&User{},
		&Wechat{},
		&Option{},
		&Question{},
	)
}

func DB() *gorm.DB {
	if db == nil {
		initDB()
	}
	return db
}

func Close() error {
	return DB().Close()
}
