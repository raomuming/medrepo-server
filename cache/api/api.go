package api

import (
	"fmt"
	"strconv"

	"medrepo-server/cache"
	"medrepo-server/mlog"
)

const exp = 60 * 10 // 10 second

func Get(uid uint) int {
	v, err := cache.Do("GET", getKey(uid))
	if err != nil {
		mlog.Error("get cache code error")
	}
	val, ok := v.([]byte)
	if !ok {
		return 0
	}
	n, _ := strconv.Atoi(string(val))
	return n
}

func Add(uid uint) error {
	key := getKey(uid)
	_, err := cache.Do("INCR", key)
	if err != nil {
		mlog.Error("set cache code err:", mlog.Err(err))
	}
	_, err = cache.Do("Expire", key, exp)
	if err != nil {
		mlog.Error("set cache code err:", mlog.Err(err))
	}
	return err
}

func getKey(uid uint) string {
	return fmt.Sprintf("api.count.%d", uid)
}
