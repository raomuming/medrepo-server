package middleware

import (
	"time"
	"strings"

	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"

	"medrepo-server/config"
	"medrepo-server/cache/api"
)

func jwtMiddle(ctx iris.Context) {
	if skipJWT(ctx.Path()) {
		ctx.Next()
		return
	}

	// token verify
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Get().JwtSecret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	// jwt verify
	if err := jwtHandler.CheckJWT(ctx); err != nil {
		ctx.StopExecution()
		return
	}

	// token info verify
	token := ctx.Values().Get("jwt").(*jwt.Token)
	userID, ok := token.Claims.(jwt.MapClaims)["user_id"]
	if !ok {
		ctx.JSON(map[string]interface{}{
			"status": 404,
			"msg": "用户尚未登录，获取用户信息失败",
		})
		ctx.StopExecution()
		return
	}

	// set user id
	ctx.Values().Set("user_id", userID)
	uid := uint(userID.(float64))
	if api.Get(uid) > 300 {
		ctx.JSON(map[string]interface{}{
			"status": 403,
			"msg": "访问过于频繁，休息一会儿吧",
		})
		ctx.StopExecution()
		return
	}
	api.Add(uid)
	ctx.Next()
}

func skipJWT(path string) bool {
	urls := []string{
		"/login",
		"/notices",
		"/webhook",
		"/helps",
	}
	for _, v := range urls {
		if v == path || strings.Contains(path, "debug") {
			return true
		}
	}
	return false
}

func GetUserID(ctx iris.Context) uint {
	uid := ctx.Values().Get("user_id")
	switch uid.(type) {
	case float64:
		return uint(uid.(float64))
	}
	return 0
}

func CreateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"end": time.Now().Unix() + 3600*24*15,
		"start": time.Now().Unix(),
	})

	return token.SignedString([]byte(config.Get().JwtSecret))
}
