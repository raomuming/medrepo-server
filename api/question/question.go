package question

import (
	"medrepo-server/model"
	"medrepo-server/api"
	
	"github.com/kataras/iris/v12"
	"gopkg.in/go-playground/validator.v9"
	"github.com/jinzhu/gorm"
)

// example: https://github.com/mohuishou/scuplus-go/blob/master/api/lost_find/lost_find.go

type NewParam struct {
	ID uint `json:"id"`
	Topic    string   `json:topic`
	Options  []string `json:"options"`
	Answer   int      `json:"answer"`
	Analysis string   `json:"analysis"`
	Chapters []uint `json:chapters`
}

func Create(ctx iris.Context) {
	data := param(ctx)
	if data == nil {
		return
	}
	
	res := model.DB().Create(data)
	if err := res.Error; err != nil {
		api.Error(ctx, 80001, "create failed!", err)
		return
	}

	api.Success(ctx, "create success!", nil)
	return
}

func param(ctx iris.Context) *model.Question {
	params := NewParam{}
	if err := ctx.ReadJSON(&params); err != nil {
		api.Error(ctx, 80400, "params error", err)
		return nil
	}

	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		api.Error(ctx, 80400, "params error!", err)
		return nil
	}

	return &model.Question{
		Model: gorm.Model{ID: params.ID},
		Topic: params.Topic,
		Options: params.Options,
		Answer: params.Answer,
		Analysis: params.Analysis,
		Chapters: params.Chapters,
	}
}