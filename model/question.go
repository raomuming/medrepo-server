package model

import (
	"github.com/jinzhu/gorm"
)

type Option struct {
	gorm.Model
	Order uint
	Description string
}

type Question struct {
	gorm.Model
	Topic    string   `json:topic`
	Options  []Option `json:"options"`
	Answer   uint      `json:"answer"`
	Analysis string   `json:"analysis"`
}