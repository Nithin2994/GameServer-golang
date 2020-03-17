package main

import (
	"encoding/json"
	"net/http"
)

func updatePlayerLevel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentType)

	var request updatePlayerLevelRequest
	var response updatePlayerLevelResponse

	_ = json.NewDecoder(r.Body).Decode(&request)
	token := r.Header.Get("token")
	response.Status = success
	if !validateToken(request.PlayerName, token) == true {
		response.Status = invalidToken
	} else {
		response.Status = updateLevel(request.PlayerName, request.PlayerLevel)
	}

	json.NewEncoder(w).Encode(&response)
}

func getPlayerState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentType)

	var request getPlayerStateRequest
	var response getPlayerStateResponse

	_ = json.NewDecoder(r.Body).Decode(&request)
	token := r.Header.Get("token")
	response.Status = success
	if !validateToken(request.PlayerName, token) == true {
		response.Status = invalidToken
	} else {
		response.PlayerName = request.PlayerName
		response.PlayerLevel = getPlayerLevel(request.PlayerName)
		response.PlayerWallet = getBalances(request.PlayerName)
		response.Status = success
	}

	json.NewEncoder(w).Encode(&response)
}
