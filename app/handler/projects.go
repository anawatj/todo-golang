package handler

import (
	"encoding/json"
	"net/http"

	"../model"
	"github.com/jinzhu/gorm"
)

func GetAllProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	projects := []model.Project{}
	db.Find(&projects)
	respondJSON(w, http.StatusOK, projects)
}

func CreateProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
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
