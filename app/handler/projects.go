package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"../model"
	"github.com/jinzhu/gorm"
)

func GetAllProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	authorizeToken := r.Header.Get("Authorization")
	err := validToken(authorizeToken, db)
	if err != nil {
		respondError(w, http.StatusUnauthorized, "UnAuthorize")
		return
	}
	projects := []model.Project{}
	db.Find(&projects)
	if len(projects) == 0 {
		respondError(w, http.StatusNotFound, "record not found")
		return
	}
	respondJSON(w, http.StatusOK, projects)
}

func CreateProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	authorizeToken := r.Header.Get("Authorization")
	err := validToken(authorizeToken, db)
	if err != nil {
		respondError(w, http.StatusUnauthorized, "UnAuthorize")
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
func UpdateProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	authorizeToken := r.Header.Get("Authorization")
	err := validToken(authorizeToken, db)
	if err != nil {
		respondError(w, http.StatusUnauthorized, "UnAuthorize")
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
	respondJSON(w, http.StatusOK, project)
}
func GetProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	authorizeToken := r.Header.Get("Authorization")
	err := validToken(authorizeToken, db)
	if err != nil {
		respondError(w, http.StatusUnauthorized, "UnAuthorize")
		return
	}
	vars := mux.Vars(r)
	name := vars["name"]

	project := model.Project{}
	if err := db.First(&project, model.Project{Name: name}).Error; err != nil {
		if err.Error() == "record not found" {
			respondError(w, http.StatusNotFound, err.Error())
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return

	}
	respondJSON(w, http.StatusOK, project)

}
