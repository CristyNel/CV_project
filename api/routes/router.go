package routes

import (
	"net/http"

	"github.com/CristyNel/CV_project/tree/main/api/handlers"
	"github.com/CristyNel/CV_project/tree/main/api/internal/app"
	"github.com/CristyNel/CV_project/tree/main/api/internal/auth"
	"github.com/gorilla/mux"
)

// InitializeRouter initializes the router and returns it
func InitializeRouter(app *app.App) *mux.Router {
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.Home(app, w, r)
	}).Methods("GET")

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		auth.LoginHandler(app, w, r)
	}).Methods("POST")

	r.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		auth.SignupHandler(app, w, r)
	}).Methods("POST")

	// Protected routes
	r.Handle("/users", app.RequireAuthentication(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.ShowUsers(app, w, r)
	}))).Methods("GET")

	r.Handle("/user", app.RequireAuthentication(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUser(app, w, r)
	}))).Methods("POST")

	r.Handle("/user/{id}", app.RequireAuthentication(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.ShowUser(app, w, r)
	}))).Methods("GET")

	r.Handle("/user/{id}", app.RequireAuthentication(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateUser(app, w, r)
	}))).Methods("PUT")

	r.Handle("/user/{id}", app.RequireAuthentication(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteUser(app, w, r)
	}))).Methods("DELETE")

	r.Handle("/pdf", app.RequireAuthentication(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.GenerateTemplate(app, w, r)
	}))).Methods("GET")

	r.Handle("/logout", app.RequireAuthentication(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth.LogoutHandler(app, w, r)
	}))).Methods("POST")

	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API is running"))
	}).Methods("GET")

	return r
}
