package handler

import (
	"net/http"

	"../model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

func GetAllTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	claims, err := verifyToken(r)
	if err != nil {
		respondError(w, http.StatusUnauthorized, "Unauthorize")
		return
	}
	mapClaim := claims.(jwt.MapClaims)
	auth := model.Auth{
		UserName: mapClaim["UserName"].(string),
		Password: mapClaim["Password"].(string),
	}

	user := model.User{}
	if err := db.First(&user, model.User{UserName: auth.UserName}).Error; err != nil {
		respondError(w, http.StatusUnauthorized, "Unauthorize")
		return
	}
	tasks := []model.Task{}
	db.Find(&tasks)
	respondJSON(w, http.StatusOK, tasks)

}
