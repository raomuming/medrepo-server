package model

import (
	"github.com/jinzhu/gorm"
)

type Option struct {
	gorm.Model
	Order uint
	Description string `gorm:"type:text"`
}

type Question struct {
	gorm.Model
	Topic    string   `gorm:"type:text" json:topic`
	Options  []Option `json:"options"`
	Answer   uint      `json:"answer"`
	Analysis string   `gorm:"type:text" json:"analysis"`
}