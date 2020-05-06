package wechat

import (
	"log"
	"net/http"
	"errors"
	"fmt"
	"encoding/json"
	"io/ioutil"

	"medrepo-server/cache/token"
	"medrepo-server/config"
)

// Token wechat access token
type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int `json:"expires_in`
}

func GetAccessToken(refresh bool) (string, error) {
	t := token.Get()
	if !refresh && t != "" {
		return t, nil
	}

	// 从微信服务器获取token
	c := http.Client{}
	resp, err := c.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		config.Get().Wechat.Appid,
		config.Get().Wechat.Secret))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	at := Token{}
	json.Unmarshal(body, &at)

	// cache token
	if at.AccessToken == "" {
		log.Println(at)
		return "", errors.New("获取token失败!")
	}
	err = token.Set(at.AccessToken)
	if err != nil {
		return "", err
	}
	return at.AccessToken, nil
}