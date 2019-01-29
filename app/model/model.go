package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Project struct {
	gorm.Model
	Name        string `gorm:"unique" json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type Task struct {
	gorm.Model
	Name        string `gorm:"unique" json:"name"`
	Description string `json:"description"`
	Subject     string `json:"subject"`
	ProjectID   int    `json:"projectId"`
}

type User struct {
	gorm.Model
	UserName  string `gorm:"unique" json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type Auth struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
type JwtToken struct {
	Token string `json:"token"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Project{}, Task{}, User{})
	return db
}
