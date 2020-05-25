package api

import (
	"medrepo-server/mlog"
	"medrepo-server/model"
	
	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	mlog.Info("route", mlog.String("path", "/login"))
	code := ctx.FormValue("code")

	if code == "" {
		mlog.Error("route-login", mlog.String("error", "code is empty"))
		Error(ctx, 10400, "code can not be empty", nil)
		return
	}

	// obtain open id
	user := model.User{}
	if err := user.Wechat.GetOpenid(code); err != nil {
		mlog.Error("route-login", mlog.String("error", "failed to get user info"))
		Error(ctx, 10401, "failed to get user info", nil)
		return
	}

	// login
	token, err := user.Login()
	if err != nil {
		mlog.Error("user.Login() failed", mlog.Err(err))
		Error(ctx, 10401, "login failed", nil)
		return
	}

	Success(ctx, "login success!", map[string]interface{}{
		"token": token,
	})
}