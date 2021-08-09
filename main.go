package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Nithin2994/gomodules/redis"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var contentType string = "application/json"
var success string = "Success"
var failed string = "Failed"
var waiting string = "Waiting"
var invalidToken string = "Invalid Token"

func packagesExample() {
	fmt.Println(redis.GetRedisConnection())
}

func fileExample() {
	fileOp()
}
func main() {
	// packagesExample()
	// fileOp()
	service()
}

func service() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},           // All origins
		AllowedMethods: []string{"GET", "POST"}, // Allowing only get, just an example
	})

	router := mux.NewRouter()
	router.HandleFunc("/index", index)
	router.HandleFunc("/login", login).Methods("POST", "OPTIONS")
	router.HandleFunc("/register", register).Methods("POST", "OPTIONS")
	router.HandleFunc("/leaderboard/getActiveLeaderboard", getActiveLeaderboard).Methods("GET", "OPTIONS")
	router.HandleFunc("/leaderboard/{leaderboardId}/addScore", addScore).Methods("POST", "OPTIONS")
	router.HandleFunc("/leaderboard/create", create).Methods("POST", "OPTIONS")
	router.HandleFunc("/leaderboard/delete", delete).Methods("POST", "OPTIONS")
	router.HandleFunc("/leaderboard/{leaderboardId}", getLeaderboard).Methods("GET", "OPTIONS")
	router.HandleFunc("/leaderboard/{leaderboardId}/{playerId}", playerDetails).Methods("GET", "OPTIONS")
	router.HandleFunc("/matchmaking/findOpponent", findOpponenet).Methods("POST", "OPTIONS")
	router.HandleFunc("/matchmaking/pollOpponent", pollOpponent).Methods("POST", "OPTIONS")
	router.HandleFunc("/wallet/updateWalletBalances", updatePlayerWallet).Methods("POST", "OPTIONS")
	router.HandleFunc("/wallet/balances", getPlayerWallet).Methods("POST", "OPTIONS")
	router.HandleFunc("/wallet/debit", debit).Methods("POST", "OPTIONS")
	router.HandleFunc("/wallet/credit", credit).Methods("POST", "OPTIONS")
	router.HandleFunc("/player/updatePlayerLevel", updatePlayerLevel).Methods("POST", "OPTIONS")
	router.HandleFunc("/player/getPlayerState", getPlayerState).Methods("POST", "OPTIONS")
	log.Fatal(http.ListenAndServe(":8081", c.Handler(router)))

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Welcome to Game Server</h1>")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "<h2>Below are the apis</h2>")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "<h3>Login Apis</h2>")
	fmt.Fprintln(w, "<ul>")
	fmt.Fprintln(w, "<li>POST		/login</li>")
	fmt.Fprintln(w, "<li>POST		/register</li>")
	fmt.Fprintln(w, "</ul>")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "<h3>Leaderboard Apis</h2>")
	fmt.Fprintln(w, "<ul>")
	fmt.Fprintln(w, "<li>POST		/leaderboard/create</li>")
	fmt.Fprintln(w, "<li>POST		/leaderboard/delete</li>")
	fmt.Fprintln(w, "<li>GET		/leaderboard/getActiveLeaderboard</li>")
	fmt.Fprintln(w, "<li>GET		/leaderboard/{leaderboardId}</li>")
	fmt.Fprintln(w, "<li>POST		/leaderboard/{leaderboardId}/addScore</li>")
	fmt.Fprintln(w, "<li>GET		/leaderboard/{leaderboardId}/{playerId}</li>")
	fmt.Fprintln(w, "</ul>")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "<h3>Matchmaking Apis</h2>")
	fmt.Fprintln(w, "<ul>")
	fmt.Fprintln(w, "<li>POST		/matchmaking/findOpponent</li>")
	fmt.Fprintln(w, "<li>POST		/matchmaking/pollOpponent</li>")
	fmt.Fprintln(w, "</ul>")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "<h3>Wallet Apis</h2>")
	fmt.Fprintln(w, "<ul>")
	fmt.Fprintln(w, "<li>POST		/wallet/updateWalletBalances</li>")
	fmt.Fprintln(w, "<li>POST		/wallet/balances</li>")
	fmt.Fprintln(w, "<li>POST		/wallet/debit</li>")
	fmt.Fprintln(w, "<li>POST		/wallet/credit</li>")
	fmt.Fprintln(w, "</ul>")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "<h3>Player Apis</h2>")
	fmt.Fprintln(w, "<ul>")
	fmt.Fprintln(w, "<li>POST		/player/getPlayerState</li>")
	fmt.Fprintln(w, "<li>POST		/player/updatePlayerLevel</li>")

}
