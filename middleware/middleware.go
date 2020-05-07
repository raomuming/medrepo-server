package middleware

import (
	"github.com/kataras/iris/v12"
)

func Register(app *iris.Application) {
	app.Use(jwtMiddle)
}