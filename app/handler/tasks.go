package handler

import (
	"encoding/json"
	"net/http"

	"../model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	authorizeToken := r.Header.Get("Authorization")
	err := validToken(authorizeToken, db)
	if err != nil {
		respondError(w, http.StatusUnauthorized, "UnAuthorize")
		return
	}

	vars := mux.Vars(r)
	name := vars["name"]
	tasks := []model.Task{}
	db.Table("tasks").Select("tasks.id,tasks.name,tasks.description,tasks.project_id").Joins("join projects on tasks.project_id=projects.id").Where("projects.name=?", name).Scan(&tasks)
	if len(tasks) == 0 {
		respondError(w, http.StatusNotFound, "record not found")
		return
	}
	respondJSON(w, http.StatusOK, tasks)

}

func CreateTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	authorizeToken := r.Header.Get("Authorization")
	err := validToken(authorizeToken, db)
	if err != nil {
		respondError(w, http.StatusUnauthorized, "UnAuthorize")
		return
	}

	task := model.Task{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&task).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, task)
}
