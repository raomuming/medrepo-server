package middleware

import (
	"time"
	"strings"

	"medrepo-server/config"
	"medrepo-server/cache/api"

	"github.com/kataras/iris/v12"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

func jwtMiddle(ctx iris.Context) {
	if (skipJWT(ctx.Path())) {
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
			"status": 401,
			"msg": "user is not login, get user info failed",
		})
		ctx.StopExecution()
		return
	}

	// token expiration check

	// set user id
	ctx.Values().Set("user_id", userID)
	uid := uint(userID.(float64))
	if api.Get(uid) > 300 {
		ctx.JSON(map[string]interface{}{
			"status": 403,
			"msg": "request limits exceed",
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
		"/helps",
		"/webhook",
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