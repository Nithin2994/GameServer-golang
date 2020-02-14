package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

var expireTime int

func getCurrenctLeaderboard() string {
	conn, err := getRedisConnection()
	defer conn.Close()

	id, err := redis.String(conn.Do("GET", "currentLeaderboard"))
	if err != nil {
		return ""
	}
	return id
}

func getRandomLeaderBoardID() string {
	conn, err := getRedisConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	name := getCurrenctLeaderboard()
	defer conn.Close()
	if name != "" {
		ids, err := redis.Strings(conn.Do("SRANDMEMBER", name, 1))
		if err != nil {
			return ""
		}
		if ids != nil {
			return ids[0]
		}
	}
	return ""
}

func deleteLeaderboard(name string) {
	conn, err := getRedisConnection()
	defer conn.Close()

	ids, err := redis.Strings(conn.Do("SMEMBERS", name))
	if err != nil {
		log.Fatal(err)
	}
	for _, id := range ids {
		fmt.Println("DEL " + id)
		conn.Do("DEL", id)
	}
	conn.Do("DEL", "currentLeaderboard", name)
}

func createLeaderboard(l leaderboard, id string) {
	conn, err := getRedisConnection()
	defer conn.Close()

	expireTime = l.Validity * 60
	i := 1
	for i <= 3 {
		newID := id + "_" + strconv.FormatInt(int64(i), 10)
		conn.Do("SADD", l.Name, newID)
		conn.Do("ZADD", newID)
		conn.Do("EXPIRE", newID, expireTime)
		i = i + 1
	}
	_, err = conn.Do("SET", "currentLeaderboard", l.Name)
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Do("SET", l.ID, l)
	if err != nil {
		log.Fatal(err)
	}

	conn.Do("EXPIRE", "currentLeaderboard", expireTime)
	conn.Do("EXPIRE", l.ID, expireTime)
	conn.Do("EXPIRE", l.Name, expireTime)
}

func addToLeaderboard(s int, n string, name string) int {
	conn, err := getRedisConnection()
	defer conn.Close()

	id, err := redis.String(conn.Do("GET", n))
	if err != nil {
		log.Fatal(err)
	}
	if id == "" {
		id = getRandomLeaderBoardID()
		if id != "" {
			fmt.Println("id : " + id)
			fmt.Println("n : " + n)
			_, err = conn.Do("ZADD", id, s, n)
			_, err = conn.Do("SET", n, id)
			expTime, err := redis.Int(conn.Do("PTTL", id))
			conn.Do("EXPIRE", n, expTime/1000)
			if err != nil {
				return -1
			}
		} else {
			return -1
		}
	} else {
		isMember, err := redis.Bool(conn.Do("SISMEMBER", name, id))
		if isMember {
			_, err = conn.Do("ZINCRBY", id, s, n)
			if err != nil {
				return -1
			}
		} else {
			conn.Do("DEL", n)
			addToLeaderboard(s, n, name)
		}
	}
	return 0
}

func getLeaderboardDetails(id string) map[string]scores {
	conn, err := getRedisConnection()
	defer conn.Close()

	detailsMap := make(map[string]scores)

	leaderboards, err := redis.Strings(conn.Do("SMEMBERS", id))
	if err != nil {
		log.Fatal(err)
	}
	for _, lb := range leaderboards {
		details, _ := redis.StringMap(conn.Do("ZREVRANGEBYSCORE", lb, "inf", "-inf", "withscores"))
		detailsMap[lb] = details
		fmt.Println(details)
	}

	fmt.Println(detailsMap)
	return detailsMap
}

func getLeaderboardPositions(id string) []string {
	conn, err := getRedisConnection()
	defer conn.Close()

	details, err := redis.Strings(conn.Do("ZREVRANGEBYSCORE", id, "inf", "-inf"))
	if err != nil {
		log.Fatal(err)
	}
	return details
}

func getPlayerRank(playerID string) int {
	conn, err := getRedisConnection()
	defer conn.Close()

	var playerRank int

	leaderboardID, err := redis.String(conn.Do("GET", playerID))
	if err != nil {
		log.Fatal(err)
	}
	if leaderboardID != "" {
		rank, err := redis.Int(conn.Do("ZREVRANK", leaderboardID, playerID))
		if err != nil {
			log.Fatal(err)
		}
		playerRank = rank
	}

	return playerRank + 1
}

func getPlayerScore(playerID string) int {
	conn, err := getRedisConnection()
	defer conn.Close()
	var playerScore int
	leaderboardID, err := redis.String(conn.Do("GET", playerID))
	if err != nil {
		return -1
	}
	if leaderboardID != "" {
		score, err := redis.Int(conn.Do("ZSCORE", leaderboardID, playerID))
		if err != nil {
			return -1
		}
		playerScore = score
	}
	return playerScore
}

func getTopNPlayers(playerID string) []string {
	conn, err := getRedisConnection()
	defer conn.Close()

	leaderboardID, err := redis.String(conn.Do("GET", playerID))
	if err != nil {
		log.Fatal(err)
	}
	if leaderboardID != "" {
		details, err := redis.Strings(conn.Do("ZREVRANGEBYSCORE", leaderboardID, "inf", "-inf", "LIMIT", 0, 10))
		if err != nil {
			log.Fatal(err)
		}
		return details
	}
	return nil
}

func getAboveAndBelow(playerID string) []string {
	conn, err := getRedisConnection()
	defer conn.Close()

	leaderboardID, err := redis.String(conn.Do("GET", playerID))
	playerRank := getPlayerRank(playerID)
	if err != nil {
		log.Fatal(err)
	}
	if leaderboardID != "" {
		var beginIndex int
		var endIndex int
		if playerRank-4 < 0 {
			beginIndex = 0
		} else {
			beginIndex = playerRank - 4
		}
		endIndex = playerRank + 4

		above, err := redis.Strings(conn.Do("ZREVRANGEBYSCORE", leaderboardID, "inf", "-inf", "LIMIT", beginIndex, playerRank))
		fmt.Println(above)
		below, err := redis.Strings(conn.Do("ZREVRANGEBYSCORE", leaderboardID, "inf", "-inf", "LIMIT", playerRank, endIndex))
		fmt.Println(below)
		if err != nil {
			log.Fatal(err)
		}
		return append(above, below...)
	}
	return nil
}
