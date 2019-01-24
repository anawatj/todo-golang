package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../../config"
	"../model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		respondError(w, http.StatusUnauthorized, err.Error())
		return
	}
	user.Password = string(hash[:])
	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, "Sign up success")
}

func SignIn(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	auth := model.Auth{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&auth); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	user := model.User{}
	if err := db.First(&user, model.User{UserName: auth.UserName}).Error; err != nil {
		respondError(w, http.StatusUnauthorized, "user name or password are in correct")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(auth.Password))
	if err != nil {
		respondError(w, http.StatusUnauthorized, "Username or password are incorrect")
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserName": auth.UserName,
		"Password": auth.Password,
	})
	tokenString, error := token.SignedString([]byte(config.GetJwtKey().Key))
	if error != nil {
		fmt.Println(error)
	}
	jwtToken := model.JwtToken{}
	jwtToken.Token = tokenString
	respondJSON(w, http.StatusOK, jwtToken)
}
