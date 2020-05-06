package middleware

import (
	"github.com/kataras/iris/v12"
)

// register jwt middle
func Register(app *iris.Application) {
	app.Use(jwtMiddle)
}
