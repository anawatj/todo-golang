package handler

import (
	"net/http"

	"../model"
	"github.com/jinzhu/gorm"
)

func GetAllTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	tasks := []model.Task{}
	db.Find(&tasks)
	respondJSON(w, http.StatusOK, tasks)

}
