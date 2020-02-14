package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func register(w http.ResponseWriter, r *http.Request) {
	var res response
	var req registrationRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	status := savePlayer(req.PlayerName, req.Password, req.Email, req.MobileNo)
	if status == 0 {
		res.Status = success
	} else {
		res.Status = failed
	}
	createSession(req.PlayerName)
	json.NewEncoder(w).Encode(&res)
}

func login(w http.ResponseWriter, r *http.Request) {
	var res loginResponse
	var req loginRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	status := checkPlayer(req.PlayerName, req.Password)
	fmt.Println("status ", status)
	if status == 0 {
		res.Status = success
		token := createSession(req.PlayerName)
		res.Token = token
	} else {
		res.Status = failed
	}
	json.NewEncoder(w).Encode(&res)
}

var jwtKey = []byte("111222")

func createSession(player string) string {

	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &claims{
		Username: player,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	sessionKey := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := sessionKey.SignedString(jwtKey)

	if err != nil {
		return "Error"
	}

	logSessionToken(player, tokenString)
	return tokenString
}
