package model

import (
	"github.com/jinzhu/gorm"
)

type Course struct {
	gorm.Model
	Name     string
	Description string
	Chapters []Chapter
}
