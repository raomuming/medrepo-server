package api

import (
	//"fmt"
	//"time"

	"medrepo-server/mlog"
	"medrepo-server/model"

	"github.com/kataras/iris/v12"
)

func GetQuestion(ctx iris.Context) {
	mlog.Info("route", mlog.String("path", "/question"))
	id, _ := ctx.Params().GetInt("id")
	//id, err := ctx.Params().GetInt("id")
	/*
	if err != nil || id == 0 {
		mlog.Error("route-getQuestion", mlog.Err(err))
		Error(ctx, 50400, "parameter error", nil)
		return
	}*/
	question := model.Question{}
	model.DB().Find(&question, id)
	Success(ctx, "get question successed", question)
}