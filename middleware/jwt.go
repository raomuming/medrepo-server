package middleware

import (
	"github.com/kataras/iris/v12"
)

func jwtMiddle(ctx iris.Context) {
	ctx.Next()
}