package model

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"medrepo-server/config"
)

var db *gorm.DB

type Model struct {
	ID uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func initDB() {
	conf := config.Get().Mysql

	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.DB))
	if err != nil {
		log.Fatal("数据库连接错误:", err, conf)
	}
	db.DB().SetMaxOpenConns(20)
	db.DB().SetConnMaxLifetime(10 * time.Second)
	db.LogMode(config.Get().Debug)

	autoMigrate()
}

func autoMigrate() {
	db.AutoMigrate(
		&User{}
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
