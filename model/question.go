package model

import (
	"github.com/jinzhu/gorm"
)

type Option struct {
	gorm.Model
	Order uint
	Description string `gorm:"type:varchar(1023)"`
	QuestionID uint
}

type Question struct {
	gorm.Model
	Topic    string   `gorm:"type:varchar(1023)" json:topic`
	Options  []Option `json:"options"`
	Answer   uint      `json:"answer"`
	Analysis string   `gorm:"type:varchar(2047)" json:"analysis"`
}