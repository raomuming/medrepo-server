package route

import (
	"github.com/kataras/iris/v12"

	"medrepo-server/api"
	"medrepo-server/api/question"
)

func Routes(app *iris.Application) {
	app.Post("/login", api.Login)
	app.Get("/question", question.Get)
	app.Post("/question", question.Create)
}