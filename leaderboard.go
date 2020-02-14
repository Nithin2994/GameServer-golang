package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getActiveLeaderboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentType)

	var response createLeaderboardResponse

	id := getCurrenctLeaderboard()

	response.ID = id
	response.Status = success

	json.NewEncoder(w).Encode(&response)
}

func playerDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentType)

	var response playerLeaderboardDetailsResponse

	params := mux.Vars(r)
	playerID := params["playerId"]
	token := r.Header.Get("token")

	if !validateToken(playerID, token) == true {
		response.Status = invalidToken
		response.StatusCode = -3
	} else {
		rank := getPlayerRank(playerID)
		topN := getTopNPlayers(playerID)
		aboveAndBelow := getAboveAndBelow(playerID)
		playerScore := getPlayerScore(playerID)

		response.Rank = rank
		response.Top = topN
		response.Score = playerScore
		response.AboveAndBelow = aboveAndBelow
		response.Status = success
		response.StatusCode = 0
	}

	json.NewEncoder(w).Encode(&response)
}

func addScore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentType)

	var s score
	var res response

	_ = json.NewDecoder(r.Body).Decode(&s)
	token := r.Header.Get("token")

	if !validateToken(s.Name, token) == true {
		res.Status = invalidToken
		res.StatusCode = -3
	} else {
		params := mux.Vars(r)
		leaderboardNameReq := params["leaderboardId"]
		leaderboardName := getCurrenctLeaderboard()
		if leaderboardName != "" && leaderboardNameReq == leaderboardName {
			status := addToLeaderboard(s.Score, s.Name, leaderboardName)
			if status == -1 {
				res.Status = "Error While Updating Leaderboard"
				res.StatusCode = -1
			} else {
				res.Status = success
				res.StatusCode = 0
			}
		} else {
			res.Status = "Invalid Leaderboard"
			res.StatusCode = -2
		}
	}
	json.NewEncoder(w).Encode(&res)
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentType)
	var name leaderboard
	_ = json.NewDecoder(r.Body).Decode(&name)

	fmt.Println(name.Name)

	deleteLeaderboard(name.Name)

	var res response
	res.Status = success
	res.StatusCode = 0

	json.NewEncoder(w).Encode(&res)
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentType)

	var response createLeaderboardResponse
	var l leaderboard
	var id string

	_ = json.NewDecoder(r.Body).Decode(&l)

	id = "Leaderboard_" + l.Name
	l.ID = id
	createLeaderboard(l, id)

	response.ID = l.Name
	response.Status = success

	json.NewEncoder(w).Encode(&response)
}

func getLeaderboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentType)

	var response leaderboardDetailsResponse

	params := mux.Vars(r)
	id := params["leaderboardId"]

	res2 := getLeaderboardDetails(id)

	response.Details = res2

	fmt.Println(response)
	json.NewEncoder(w).Encode(&response)
}
