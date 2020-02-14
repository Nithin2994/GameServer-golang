package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var randomMatchKey string = "RANDOM_MATCH"

func getPlayerFromQueue(player string) string {
	conn, err := getRedisConnection()
	defer conn.Close()

	opp, _ := redis.String(conn.Do("GET", "match_"+player))

	if opp != "" {
		return opp
	}

	size, err := redis.Int(conn.Do("LLEN", randomMatchKey))
	if err != nil {
		return "Error"
	}
	if size > 0 {
		for i := 0; i < size; i++ {
			opponent, err := redis.String(conn.Do("LPOP", randomMatchKey))
			isOnline, err := redis.String(conn.Do("GET", "PLAYER_ONLINE_"+opponent))
			fmt.Println("opponent :" + opponent + ", isOnline :" + string(isOnline))
			if opponent != "" && isOnline == "1" && opponent != player {
				_, err = conn.Do("SET", "match_"+player, opponent)
				_, err = conn.Do("SET", "match_"+opponent, player)
				_, err = conn.Do("EXPIRE", "match_"+opponent, 30)
				_, err = conn.Do("EXPIRE", "match_"+player, 30)
				if err != nil {
					return "Error"
				}
				return opponent
			}
		}
	}

	_, err = conn.Do("RPUSH", randomMatchKey, player)
	_, err = conn.Do("SET", "PLAYER_ONLINE_"+player, 1)
	_, err = conn.Do("EXPIRE", "PLAYER_ONLINE_"+player, 30)
	if err != nil {
		return "Error"
	}

	return "No Match Found"
}

func getMatch(player string) string {
	conn, err := getRedisConnection()
	defer conn.Close()
	if err != nil {
		return "Error"
	}
	opp, err := redis.String(conn.Do("GET", "match_"+player))

	if opp != "" {
		return opp
	}

	return "No Match Found"
}
