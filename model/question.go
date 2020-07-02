package model

import (
	//"strings"
	//"github.com/jinzhu/gorm"
)

type Question struct {
	Model
	Topic string `json:topic`
	Options []string `json:"options"`
	Answer int `json:"answer"`
	Analysis string `json:"analysis"`
}