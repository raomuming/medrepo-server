package api

import (
	"medrepo-server/mlog"
)

import (
	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	mlog.Info("route", mlog.String("path", "/login"))
}