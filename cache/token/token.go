package token

import (
	"log"

	"medrepo-server/cache"
)

const key = "wechatAccessToken"
const expireTime = 3600 * 1.5

// Get get
func Get() string {
	v, err := cache.Do("GET", key)
	if err != nil {
		log.Println("get cache token err:", err)
	}
	t, ok := v.([]byte)
	if !ok {
		return ""
	}
	return string(t)
}

func Set(v string) error {
	_, err := cache.Do("SET", key, v)
	if err != nil {
		log.Println("set cache token err:", err)
	}
	_, err = cache.Do("Expire", key, expireTime)
	if err != nil {
		log.Println("set cache code err:", err)
	}
	return err
}
