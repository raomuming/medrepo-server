package question

import (
	"medrepo-server/model"
	"medrepo-server/api"

	"github.com/kataras/iris/v12"
	//"github.com/jinzhu/gorm"
)

func CreateQuestionList(ctx iris.Context) {
	data := createParam(ctx)
	if data == nil {
		api.Error(ctx, 80001, "create failed", nil)
		return
	}

	res := model.DB().Create(data)
	if err := res.Error; err != nil {
		api.Error(ctx, 80001, "create failed", err)
		return
	}

	api.Success(ctx, "create success", nil)
}

func createParam(ctx iris.Context) *model.QuestionList {
	return &model.QuestionList {

	}
}