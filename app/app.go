package app

import (
	"fmt"
	"log"
	"net/http"

	"../config"
	"./handler"
	"./model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize Config
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// All Route
func (a *App) setRouters() {
	a.Get("/api/v1/projects", a.GetAllProject)
	a.Post("/api/v1/projects", a.CreateProject)
	a.Get("/api/v1/tasks", a.GetAllTask)
	a.Post("/api/v1/signup", a.SignUp)
	a.Post("/api/v1/signin", a.SignIn)

}

// Route Get Method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Route Post Method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Route Put Method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Route Delete Method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// /api/v1/projects get
func (a *App) GetAllProject(w http.ResponseWriter, r *http.Request) {
	handler.GetAllProject(a.DB, w, r)
}

// /api/v1/tasks get
func (a *App) GetAllTask(w http.ResponseWriter, r *http.Request) {
	handler.GetAllTask(a.DB, w, r)
}

// /api/v1/projects post
func (a *App) CreateProject(w http.ResponseWriter, r *http.Request) {
	handler.CreateProject(a.DB, w, r)
}

// /api/v1/signup post
func (a *App) SignUp(w http.ResponseWriter, r *http.Request) {
	handler.SignUp(a.DB, w, r)
}

// /api/v1/signin post
func (a *App) SignIn(w http.ResponseWriter, r *http.Request) {
	handler.SignIn(a.DB, w, r)
}

// Run Application
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
