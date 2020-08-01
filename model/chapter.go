package model

import (
	"github.com/jinzhu/gorm"
)

type Chapter struct {
	gorm.Model
	Name     string `gorm:"type:varchar(127)" json:"name"`
	CourseID uint
}