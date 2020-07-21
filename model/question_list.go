package model

import (
	"github.com/jinzhu/gorm"
)


type QuestionList struct {
	gorm.Model
	QuestionLists []CourseQuestionList
}