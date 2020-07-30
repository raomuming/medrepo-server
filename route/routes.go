package route

import (
	"github.com/kataras/iris/v12"

	"medrepo-server/api"
	"medrepo-server/api/question"
	"medrepo-server/api/course"
)

func Routes(app *iris.Application) {
	app.Post("/login", api.Login)
	app.Get("/question", question.Get)
	app.Post("/question", question.Create)
	app.Post("/course", course.Create)
	app.Get("/course", course.Get)
	app.Put("/course/{id:int}", course.AddChapter)
}