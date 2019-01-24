package handler

import (
	"net/http"

	"../model"
	"github.com/jinzhu/gorm"
)

func GetAllTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	authorizeToken := r.Header.Get("Authorization")
	err := validToken(authorizeToken, db)
	if err != nil {
		respondError(w, http.StatusUnauthorized, "UnAuthorize")
		return
	}
	tasks := []model.Task{}
	db.Find(&tasks)
	if len(tasks) == 0 {
		respondError(w, http.StatusNotFound, "record not found")
		return
	}
	respondJSON(w, http.StatusOK, tasks)

}
