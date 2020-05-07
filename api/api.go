package api

import (
	"github.com/kataras/iris/v12"
)

func Error(ctx iris.Context, code int, msg string, data interface{}) {
	ctx.JSON(map[string]interface{}{
		"status": code,
		"msg": msg,
		"data": data,
	})
}

func Success(ctx iris.Context, msg string, data interface{}) {
	ctx.JSON(map[string]interface{}{
		"status": 0,
		"msg": msg,
		"data": data,
	})
}