package question

import (
	"medrepo-server/model"
	"medrepo-server/api"
	
	"github.com/kataras/iris/v12"
	"gopkg.in/go-playground/validator.v9"
	"github.com/jinzhu/gorm"
)

// example: https://github.com/mohuishou/scuplus-go/blob/master/api/lost_find/lost_find.go

type OptionParam struct {
	Order uint `json:order`
	Description string `json:description`
}

type NewParam struct {
	ID uint `json:"id"`
	Topic    string   `json:topic`
	Options  []OptionParam `json:"options"`
	Answer   uint      `json:"answer"`
	Analysis string   `json:"analysis"`
	Chapters []uint `json:chapters`
}

func Create(ctx iris.Context) {
	data := param(ctx)
	if data == nil {
		api.Error(ctx, 80001, "create failed!", nil)
		return
	}
	
	res := model.DB().Create(data)
	if err := res.Error; err != nil {
		api.Error(ctx, 80001, "create failed!", err)
		return
	}

	api.Success(ctx, "create success!", nil)
}

func Get(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	if id == 0 || err != nil {
		api.Error(ctx, 70400, "params error", err)
		return
	}

	question := model.Question{}
	if err := model.DB().Find(&question, id).Error; err != nil {
		api.Error(ctx, 80003, "get question failed", err)
		return
	}

	var options []model.Option	

	model.DB().Model(&question).Related(&options)
	question.Options = options

	api.Success(ctx, "success", question)
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

	answerOptions := make([]model.Option, 0)
	if options := params.Options; options != nil {
		for _, option := range options {
			answerOptions = append(answerOptions, model.Option{Order: option.Order, Description: option.Description})
		}
	}

	return &model.Question{
		Model: gorm.Model{ID: params.ID},
		Topic: params.Topic,
		Options: answerOptions,
		Answer: params.Answer,
		Analysis: params.Analysis,
	}
}