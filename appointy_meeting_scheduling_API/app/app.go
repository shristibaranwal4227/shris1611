package application

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Application has router and db instances
type Application struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize initializes the application with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
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

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/meeting", a.handleRequest(handler.GetAllProjects))
	a.Post("/meeting", a.handleRequest(handler.CreateProject))
	a.Get("/meeting/{title}", a.handleRequest(handler.GetProject))
	a.Put("/meeting/{title}", a.handleRequest(handler.UpdateProject))
	a.Delete("/meeting/{title}", a.handleRequest(handler.DeleteProject))
	a.Put("/meeting/{title}/archive", a.handleRequest(handler.ArchiveProject))
	a.Delete("/meeting/{title}/archive", a.handleRequest(handler.RestoreProject))

	// Routing for handling the tasks
	a.Get("/meeting/{title}/tasks", a.handleRequest(handler.GetAllTasks))
	a.Post("/meeting/{title}/tasks", a.handleRequest(handler.CreateTask))
	a.Get("/meeting/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.GetTask))
	a.Put("/meeting/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.UpdateTask))
	a.Delete("/meeting/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.DeleteTask))
	a.Put("/meeting/{title}/tasks/{id:[0-9]+}/complete", a.handleRequest(handler.CompleteTask))
	a.Delete("/meeting/{title}/tasks/{id:[0-9]+}/complete", a.handleRequest(handler.UndoTask))
}

// Get wraps the router for GET method
func (a *Application) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *Application) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *Application) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *Application) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *Application) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *Application) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
	
}
