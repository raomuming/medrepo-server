package model

import (
	"github.com/jinzhu/gorm"
)

// question list of a spefic course
type CourseQuestionList struct {
	gorm.Model
	CourseID uint
	Questions []Question
}