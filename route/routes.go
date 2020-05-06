package route

import (
	"github.com/kataras/iris/v12"

	"medrepo-server/api"
	"medrepo-server/api/wechat"
)

func Routes(app *iris.Application) {
	app.Post("/login", api.Login)

	app.Get("/wechat/qcode", wechat.GetQCode)
}