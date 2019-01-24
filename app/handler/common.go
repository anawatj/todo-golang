package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"../../config"
	jwt "github.com/dgrijalva/jwt-go"
)

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

func verifyToken(r *http.Request) (jwt.Claims, error) {
	authorizeToken := r.Header.Get("Authorization")
	key := config.GetJwtKey().Key
	signingKey := []byte(key)
	if len(authorizeToken) == 0 {
		err := errors.New("authorize token is required")
		return nil, err

	}
	token, err := jwt.Parse(authorizeToken, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err == nil {
		return token.Claims, nil
	} else {
		return nil, err
	}

}
