package db

import (
	"fmt"
	"goweb/utils"

	"github.com/garyburd/redigo/redis"
)

// Test db
func Test() {
	fmt.Println("I am db")
}

// ConnDB connent to redis database
func ConnDB(addr string) redis.Conn {
	r, err := redis.Dial("tcp", addr)
	if utils.ErrOccur(err) == true {
		return nil
	}
	utils.ColorPrintf("connected to database\n", utils.Green)
	return r
}

// SetDB set k-v to redis
func SetDB(r redis.Conn, key string, value string) string {
	reply, err := redis.String(r.Do("SET", key, value))
	if utils.ErrOccur(err) == true {
		return "ERROR"
	}
	return reply
}

// GetDB get k-v from redis
func GetDB(r redis.Conn, key string) string {
	value, err := redis.String(r.Do("GET", key))
	if utils.ErrOccur(err) == true {
		return ""
	}
	return value
}
