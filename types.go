package main

import "github.com/dgrijalva/jwt-go"

type createLeaderboardResponse struct {
	ID     string `json:"id"`
	Status string `json:"success"`
}

type leaderboard struct {
	Name     string `json:"name"`
	Count    int    `json:"count"`
	Validity int    `json:"validity"`
	Unit     string `json:"unit"`
	ID       string `json:"id"`
}

type score struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type response struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
}

type leaderboardDetailsResponse struct {
	Details map[string]scores `json:"details"`
}

type playerLeaderboardDetailsResponse struct {
	StatusCode    int      `json:"statusCode"`
	Status        string   `json:"status"`
	Rank          int      `json:"rank"`
	Score         int      `json:"score"`
	Top           []string `json:"top"`
	AboveAndBelow []string `json:"aboveAndBelow"`
}

type scores map[string]string

type registrationRequest struct {
	PlayerName string `json:"playerName"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	MobileNo   int64  `json:"mobileNo"`
}

type loginRequest struct {
	PlayerName string `json:"playerName"`
	Password   string `json:"password"`
}

type loginResponse struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Token      string `json:"token"`
}

type claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type player struct {
	PlayerName string `json:"playerName"`
}

type findMatchRequest struct {
	PlayerName string `json:"playerName"`
}

type findMatchResponse struct {
	Status         string `json:"status"`
	OpponentPlayer string `json:"opponentPlayer"`
}

type pollOpponentRequest struct {
	PlayerName string `json:"playerName"`
}

type pollOpponentResponse struct {
	Status         string `json:"status"`
	OpponentPlayer string `json:"opponentPlayer"`
}

type playerWallet struct {
	PlayerName string           `json:"playerName"`
	Balances   map[string]int64 `json:"balances"`
}

type updatePlayerWalletRequest struct {
	PlayerName string           `json:"playerName"`
	Balances   map[string]int64 `json:"balances"`
}

type balancesRequest struct {
	PlayerName string `json:"playerName"`
}

type creditRequest struct {
	PlayerName string `json:"playerName"`
	Amount     int64  `json:"amount"`
	Currency   string `json:"currency"`
}

type debitRequest struct {
	PlayerName string `json:"playerName"`
	Amount     int64  `json:"amount"`
	Currency   string `json:"currency"`
}

type updatePlayerWalletResponse struct {
	Status string `json:"status"`
}

type balancesResponse struct {
	Status   string           `json:"status"`
	Balances map[string]int64 `json:"balances"`
}

type creditResponse struct {
	Status string `json:"status"`
}

type debitResponse struct {
	Status string `json:"status"`
}
