package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func savePlayer(p string, pass string, email string, mobileNo int64) int {
	conn, err := getRedisConnection()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer conn.Close()

	res, _ := conn.Do("HMSET", "user_"+p, "playerName", p, "password", pass, "email", email, "mobileNo", mobileNo)
	fmt.Println(res)
	return 0
}

func checkPlayer(p string, pass string) int {
	conn, err := getRedisConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	res, _ := redis.String(conn.Do("HGET", "user_"+p, "password"))
	fmt.Println(res)
	if res == pass {
		return 0
	}
	return -1
}

func logSessionToken(p string, s string) {
	conn, err := getRedisConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	res, _ := conn.Do("SET", "session_"+p, s)
	fmt.Println(res)
}

func getSessionToken(p string) string {
	conn, err := getRedisConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	res, _ := redis.String(conn.Do("GET", "session_"+p))
	return res
}
