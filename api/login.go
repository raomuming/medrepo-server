package api

import (
	"log"

	"medrepo-server/model"
)

func Login(ctx iris.Context) {
	code := ctx.FormValue("code")
	if code == "" {
		Error(ctx, 10400, "code不能为空", nil)
		return
	}

	// get openid
	user := &model.User{}
	if err := user.Wechat.GetOpenid(code); err != nil {
		Error(ctx, 10401, "获取用户信息失败", nil)
		return
	}

	// login
	token, err := user.Login()
	if err != nil {
		log.Println(err)
		Error(ctx, 10401, "登录失败", nil)
		return
	}

	// TODO

	Success(ctx, "登录成功!", map[string]interface{}{
		"token": token,
		"verify": user.Verify,
	})
}

