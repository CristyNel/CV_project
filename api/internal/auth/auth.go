// * api/internal/auth/auth.go

package auth

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/CristyNel/CV_project/tree/main/api/internal/app"
	"github.com/CristyNel/CV_project/tree/main/api/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

// SignupRequest represents the request payload for user signup
type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginHandler handles the login requests.
func LoginHandler(app *app.App, w http.ResponseWriter, r *http.Request) {
	app.Logger.Println(" * * * ☎️  Received request for /login")

	// Check if the request method is POST. If not, return a "Method Not Allowed" error.
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data from the request. If there's an error, log it and return an "Internal Server Error".
	if err := r.ParseForm(); err != nil {
		app.Logger.Println(" * * * ⛔️ Error parsing form:", err)
		http.Error(w, " * * * Parseform, ⛔️ Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Extract the username and password from the parsed form data.
	username := strings.TrimSpace(r.Form.Get("username"))
	password := r.Form.Get("password")

	// Log the parsed username.
	app.Logger.Println("Parsed form data - Username:", username)

	// Check if the username or password is empty. If so, log it and return a "Bad Request" error.
	if username == "" || password == "" {
		app.Logger.Println("Missing username or password")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Verify the login credentials using a utility function. If valid, set the session and return a success response.
	if utils.VerifyLogin(app, username, password) {
		// Create a response struct with the username.
		response := struct {
			Username string `json:"username"`
		}{
			Username: username,
		}

		// Set the response header to indicate JSON content.
		w.Header().Set("Content-Type", "application/json")
		// Set the session for the authenticated user.
		utils.SetSession(app, username, w)
		// Write a 200 OK status to the response.
		w.WriteHeader(http.StatusOK)
		// Encode the response struct as JSON and write it to the response.
		json.NewEncoder(w).Encode(response)
	} else {
		// Log an unauthorized access attempt.
		app.Logger.Println("Unauthorized access attempt by:", username)
		// Return an "Unauthorized" error.
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

// SignupHandler handles the signup requests.
func SignupHandler(app *app.App, w http.ResponseWriter, r *http.Request) {
	app.Logger.Println(" * * * ☎️  Received request for /signup")

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if the Content-Type is application/json
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	// Decode the JSON request body
	var req SignupRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		app.Logger.Println(" * * * ⛔️ Error decoding JSON: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	email := strings.TrimSpace(req.Email)
	password := req.Password

	if email == "" || password == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = app.DB.Exec("INSERT INTO userlogin (email, password) VALUES (?,?)", email, hashedPassword)
	if err != nil {
		app.Logger.Println(" * * * ⛔️ Error querying database: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User signup successfully"))
}

// LogoutHandler handles the logout requests.
func LogoutHandler(app *app.App, w http.ResponseWriter, r *http.Request) {
	app.Logger.Println(" * * * ☎️  Received request for /logout")

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	//clear session
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("User logout successfully"))
}
