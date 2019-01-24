package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"../../config"
	"../model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
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

func verifyToken(authorizeToken string) (jwt.Claims, error) {

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
	}
	return nil, err

}
func validToken(authorizeToken string, db *gorm.DB) error {
	claims, err := verifyToken(authorizeToken)
	if err != nil {
		return err
	}
	mapClaim := claims.(jwt.MapClaims)
	auth := model.Auth{
		UserName: mapClaim["UserName"].(string),
		Password: mapClaim["Password"].(string),
	}

	user := model.User{}
	if err := db.First(&user, model.User{UserName: auth.UserName}).Error; err != nil {
		return err
	}
	return nil

}
