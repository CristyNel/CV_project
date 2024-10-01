// * CV_project/api/handlers/users.go
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/damarisnicolae/CV_project/api/internal/app"
	"github.com/damarisnicolae/CV_project/api/internal/utils"
	"github.com/damarisnicolae/CV_project/api/models"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(app *app.App, w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("Received request for /login")

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	username := strings.TrimSpace(r.Form.Get("username"))
	password := r.Form.Get("password")

	if username == "" || password == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if utils.VerifyLogin(app, username, password) {
		response := struct {
			Username string `json:"username"`
		}{
			Username: username,
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SetSession(app, username, w)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func LogoutHandler(app *app.App, w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("Received request for /logout")
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

func SignupHandler(app *app.App, w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("Received request for /signup")

	err := r.ParseForm()
	if err != nil {
		app.Logger.Println("Error parse form: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	email := strings.TrimSpace(r.Form.Get("email"))
	password := r.Form.Get("password")

	if email == "" || password == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = app.DB.Exec("INSERT INTO userlogin (username, password) VALUES (?,?)", email, hashedPassword)
	if err != nil {
		app.Logger.Println("Error querying database: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("User signup successfully"))
}

// ShowUsers handles GET requests to fetch all users
func ShowUsers(app *app.App, w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("Received request for /users, method: GET")

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := app.DB.Query("SELECT * FROM users")
	if err != nil {
		app.Logger.Println("Error querying database: ", err)
		http.Error(w, fmt.Sprintf("Error querying database: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Jobtitle, &user.Firstname, &user.Lastname, &user.Email, &user.Phone, &user.Address, &user.City, &user.Country, &user.Postalcode, &user.Dateofbirth, &user.Nationality, &user.Summary, &user.Workexperience, &user.Education, &user.Skills, &user.Languages); err != nil {
			app.Logger.Println("Error scanning row: ", err)
			http.Error(w, fmt.Sprintf("Error scanning row: %v", err), http.StatusInternalServerError)
			return
		}
		app.Logger.Printf("Scanned user: %+v\n", user)
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		app.Logger.Println("Error encoding user data: ", err)
		http.Error(w, fmt.Sprintf("Error encoding user data: %v", err), http.StatusInternalServerError)
	}
	app.Logger.Println("Successfully encoded and sent user data")
}

// CreateUser handles POST requests to create a new user
func CreateUser(app *app.App, w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("Received request for /user, method: POST")

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		app.Logger.Println("Error decoding request body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the user into the database
	_, err := app.DB.Exec("INSERT INTO users (Jobtitle, Firstname, Lastname, Email, Phone, Address, City, Country, Postalcode, Dateofbirth, Nationality, Summary, Workexperience, Education, Skills, Languages) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		user.Jobtitle, user.Firstname, user.Lastname, user.Email, user.Phone, user.Address, user.City, user.Country, user.Postalcode, user.Dateofbirth, user.Nationality, user.Summary, user.Workexperience, user.Education, user.Skills, user.Languages)
	if err != nil {
		app.Logger.Println("Error inserting user into database: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

// UpdateUser handles PUT requests to update an existing user
func UpdateUser(app *app.App, w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("Received request for /user/{id}, method: PUT")

	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract user ID from URL
	idStr := r.URL.Path[len("/user/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		app.Logger.Println("Invalid user ID: ", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		app.Logger.Println("Error decoding request body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the user in the database
	_, err = app.DB.Exec("UPDATE users SET Jobtitle = ?, Firstname = ?, Lastname = ?, Email = ?, Phone = ?, Address = ?, City = ?, Country = ?, Postalcode = ?, Dateofbirth = ?, Nationality = ?, Summary = ?, Workexperience = ?, Education = ?, Skills = ?, Languages = ? WHERE id = ?",
		user.Jobtitle, user.Firstname, user.Lastname, user.Email, user.Phone, user.Address, user.City, user.Country, user.Postalcode, user.Dateofbirth, user.Nationality, user.Summary, user.Workexperience, user.Education, user.Skills, user.Languages, id)
	if err != nil {
		app.Logger.Println("Error updating user in database: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
}

// DeleteUser handles DELETE requests to delete a user by ID
func DeleteUser(app *app.App, w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("Received request for /user/{id}, method: DELETE")

	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	userID := vars["id"]
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil || id < 1 {
		http.Error(w, "ID should be a positive integer", http.StatusBadRequest)
		return
	}

	// Check if user exists
	var exists bool
	err = app.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		app.Logger.Println("Error checking user existence: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, fmt.Sprintf("User with ID %d not found", id), http.StatusNotFound)
		return
	}

	stmt, err := app.DB.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		app.Logger.Println("Error preparing delete statement: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		app.Logger.Println("Error executing delete statement: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		app.Logger.Println("Error fetching rows affected: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, fmt.Sprintf("User with ID %d not found", id), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// AddUser handles POST requests to add a new user
func AddUser(app *app.App, w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("Received request for /user, method: POST")

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		app.Logger.Println("Error decoding request body: ", err)
		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	_, err := app.DB.Exec("INSERT INTO users (jobtitle, firstname, lastname, email, phone, address, city, country, postalcode, dateofbirth, nationality, summary, workexperience, education, skills, languages) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		user.Jobtitle, user.Firstname, user.Lastname, user.Email, user.Phone, user.Address, user.City, user.Country, user.Postalcode, user.Dateofbirth, user.Nationality, user.Summary, user.Workexperience, user.Education, user.Skills, user.Languages)
	if err != nil {
		app.Logger.Println("Error inserting user into database: ", err)
		http.Error(w, fmt.Sprintf("Error inserting user into database: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
