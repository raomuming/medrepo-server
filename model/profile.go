package model

import (
	"github.com/jinzhu/gorm"
)

// `Profile` belongs to `User`, `UserID` is the foreign key
type Profile struct {
	gorm.Model
	UserID uint
	User   User
	DefaultQuestionList []QuestionList
	UserDefinedQuestionList []QuestionList
}
