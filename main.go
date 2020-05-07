package main

import (
	"github.com/kataras/iris/v12/middleware/pprof"
	"github.com/kataras/iris/v12"

	"medrepo-server/config"
	"medrepo-server/middleware"
	"medrepo-server/route"
)

func main() {
	app := iris.New()

	middleware.Register(app)
	route.Routes(app)
	app.Any("/debug/pprof/{action:path}", pprof.New())
	app.Run(iris.Addr("0.0.0.0:" + config.Get().Port))
}