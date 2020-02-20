package main

import (
	"encoding/json"
	"net/http"
)

func updatePlayerWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentType)

	var request updatePlayerWalletRequest
	var response updatePlayerWalletResponse

	_ = json.NewDecoder(r.Body).Decode(&request)
	token := r.Header.Get("token")
	response.Status = success
	if !validateToken(request.PlayerName, token) == true {
		response.Status = invalidToken
	} else {
		response.Status = updateBalances(request.PlayerName, request.Balances)
	}

	json.NewEncoder(w).Encode(&response)
}

func getPlayerWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentType)

	var request balancesRequest
	var response balancesResponse

	_ = json.NewDecoder(r.Body).Decode(&request)
	token := r.Header.Get("token")
	response.Status = success
	if !validateToken(request.PlayerName, token) == true {
		response.Status = invalidToken
	} else {
		response.Balances = getBalances(request.PlayerName)
	}

	json.NewEncoder(w).Encode(&response)
}

func debit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentType)

	var request debitRequest
	var response debitResponse

	_ = json.NewDecoder(r.Body).Decode(&request)
	token := r.Header.Get("token")
	response.Status = success
	if !validateToken(request.PlayerName, token) == true {
		response.Status = invalidToken
	} else {
		response.Status = transaction(request.PlayerName, request.Currency, request.Amount, "DB")
	}

	json.NewEncoder(w).Encode(&response)
}

func credit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentType)

	var request creditRequest
	var response creditResponse

	_ = json.NewDecoder(r.Body).Decode(&request)
	token := r.Header.Get("token")
	response.Status = success
	if !validateToken(request.PlayerName, token) == true {
		response.Status = invalidToken
	} else {
		response.Status = transaction(request.PlayerName, request.Currency, request.Amount, "CR")
	}

	json.NewEncoder(w).Encode(&response)
}
