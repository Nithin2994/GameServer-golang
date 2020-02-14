package main

import (
	"encoding/json"
	"net/http"
)

func findOpponenet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentType)
	var response findMatchResponse
	var request findMatchRequest

	_ = json.NewDecoder(r.Body).Decode(&request)
	token := r.Header.Get("token")

	if !validateToken(request.PlayerName, token) == true {
		response.Status = invalidToken
	} else {
		opponent := getPlayerFromQueue(request.PlayerName)

		if opponent == "No Match Found" {
			response.Status = waiting
		} else if opponent == "Error" {
			response.Status = failed
		} else {
			response.Status = success
			response.OpponentPlayer = opponent
		}
	}

	json.NewEncoder(w).Encode(&response)
}

func pollOpponent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentType)

	var response pollOpponentResponse
	var request pollOpponentRequest

	_ = json.NewDecoder(r.Body).Decode(&request)
	token := r.Header.Get("token")

	if !validateToken(request.PlayerName, token) == true {
		response.Status = invalidToken
	} else {
		opponent := getMatch(request.PlayerName)

		if opponent == "No Match Found" {
			response.Status = "No Match Found"
		} else if opponent == "Error" {
			response.Status = failed
		} else {
			response.Status = success
			response.OpponentPlayer = opponent
		}
	}

	json.NewEncoder(w).Encode(&response)
}
