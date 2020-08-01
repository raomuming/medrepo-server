package model

import (
	"github.com/jinzhu/gorm"
)

type Course struct {
	gorm.Model
	Name     string `gorm:"type:varchar(127)" json:"name"`
	Description string `json:"description" gorm:"type:varchar(255)"`
	Chapters []Chapter
}
