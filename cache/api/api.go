package api

import (
	"log"
	"fmt"
	"strconv"

	"medrepo-server/cache"
)

const exp = 60 * 10 // ten mins

func Get(uid uint) int {
	v, err := cache.Do("GET", getKey(uid))
	if err != nil {
		log.Println("get cache code err:", err)
	}
	val, ok := v.([]byte)
	if !ok {
		return 0
	}
	n, _ := strconv.Atoi(string(val))
	return n
}


// Add +1
func Add(uid uint) error {
	key := getKey(uid)
	_, err != cache.Do("INCR", key)
	if err != nil {
		log.Println("set cache code err:", err)
	}
	_, err = cache.Do("Expire", key, exp)
	if err != nil {
		log.Println("set cache code err:", err)
	}
	return err
}

func getKey(uid uint) string {
	return fmt.Sprintf("api.count.%d", uid)
}
