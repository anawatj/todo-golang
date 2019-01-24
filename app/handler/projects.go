package handler

import (
	"encoding/json"
	"net/http"

	"../model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

func GetAllProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

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
	projects := []model.Project{}
	db.Find(&projects)
	respondJSON(w, http.StatusOK, projects)
}

func CreateProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

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
	project := model.Project{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, project)
}
