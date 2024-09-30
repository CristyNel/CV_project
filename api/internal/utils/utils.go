// CV_project/api/internal/utils/utils.go

package utils

import (
	"encoding/json"
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

func VerifyLogin(app *app.App, username, password string) bool {
	var hashedPassword string
	err := app.DB.QueryRow("SELECT password FROM userlogin WHERE username = ?", username).Scan(&hashedPassword)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func SetSession(app *app.App, userName string, w http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}

	if encoded, err := CookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
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
