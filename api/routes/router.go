// * CV_project/api/routes/router.go
package routes

import (
	"net/http"

	"github.com/damarisnicolae/CV_project/api/handlers"
	"github.com/damarisnicolae/CV_project/api/internal/app"

	"github.com/gorilla/mux"
)

func InitializeRouter(app *app.App) *mux.Router {
	r := mux.NewRouter()

	// routes
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.Home(app, w, r)
	}).Methods("GET")

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeUsers(app, w, r)
	}).Methods("GET")

	r.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		handlers.ShowUsers(app, w, r)
	}).Methods("GET")

	r.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUser(app, w, r)
	}).Methods("POST")

	r.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateUser(app, w, r)
	}).Methods("PUT")

	r.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteUser(app, w, r)
	}).Methods("DELETE")

	r.HandleFunc("/pdf", func(w http.ResponseWriter, r *http.Request) {
		handlers.GenerateTemplate(app, w, r)
	}).Methods("GET")

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(app, w, r)
	}).Methods("POST")

	r.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		handlers.SignupHandler(app, w, r)
	}).Methods("POST")

	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		handlers.LogoutHandler(app, w, r)
	}).Methods("POST")
	// health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API is running"))
	}).Methods("GET")

	return r
}
