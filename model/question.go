package model

import (
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	Topic    string   `json:topic`
	Options  []string `json:"options"`
	Answer   int      `json:"answer"`
	Analysis string   `json:"analysis"`
	Chapters []Chapter
}
