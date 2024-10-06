// * CV_project/api/internal/utils/utils.go
package utils

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/damarisnicolae/CV_project/api/internal/app"
	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
)

// Template struct definition
type Template struct {
	Id   int64
	Path string
}

var CookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64), // this key is used for signing
	securecookie.GenerateRandomKey(32), // this key is used for encryption
)

// * get environment variables
func GetEnv(app *app.App, key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// VerifyLogin verifies the username and password against the stored hash in the database.
func VerifyLogin(app *app.App, username, password string) bool {
	if app == nil {
		log.Println("App is nil")
		return false
	}

	if app.DB == nil {
		app.Logger.Println("Database connection is nil")
		return false
	}

	var storedHash string

	// Query the database to get the stored hash for the given username
	query := "SELECT password FROM user_session WHERE username = ?"
	err := app.DB.QueryRow(query, username).Scan(&storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			app.Logger.Println("No user found with username:", username)
			return false
		}
		app.Logger.Println("Error querying database:", err)
		return false
	}

	// Log the username and the stored hash
	app.Logger.Printf("Username found: %s", username) //, Password in DB, storedHash

	// Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		app.Logger.Println("Password does not match for user:", username)
		return false
	}

	return true
}

func SetSession(app *app.App, userName string, w http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}

	encoded, err := CookieHandler.Encode("session", value)
	if err != nil {
		app.Logger.Println("Error encoding cookie:", err) // Log encoding errors
		return
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: encoded,
		Path:  "/",
	}
	http.SetCookie(w, cookie)
	app.Logger.Println("Session set for user:", userName) // Log session creation
}

// ErrorResponse sends an error response with a given status code and message
func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// JSONResponse sends a JSON response with a given status code and payload
func JSONResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}
