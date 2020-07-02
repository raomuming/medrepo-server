package route

import (
	"github.com/kataras/iris/v12"

	"medrepo-server/api"
)

func Routes(app *iris.Application) {
	app.Post("/login", api.Login)
	app.Get("/question", api.GetQuestion)
}