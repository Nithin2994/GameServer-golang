package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func validateToken(player string, token string) bool {
	dbToken := getSessionToken(player)
	if dbToken == token {
		fmt.Println("true")
		return true
	}

	return false

}

func getRedisConnection() (redis.Conn, error) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	return conn, err
}
