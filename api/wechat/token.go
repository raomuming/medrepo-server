package wechat

import (
	"github.com/kataras/iris/v12"

	"medrepo-server/api"
	"medrepo-server/util/wechat"
)

func Token(ctx iris.Context) {
	token, err := wechat.GetAccessToken(false)
	if err != nil {
		api.Error(ctx, 400, "token获取失败", err)
		return
	}
	api.Success(ctx, "获取成功!", token)
}