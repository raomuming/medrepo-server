package model

import (
	"github.com/jinzhu/gorm"
)

type Chapter struct {
	gorm.Model
	Name     string
	CourseID uint
}
