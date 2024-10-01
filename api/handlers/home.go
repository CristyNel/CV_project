// * CV_project/api/handlers/home.go
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/damarisnicolae/CV_project/api/internal/app"
	"github.com/damarisnicolae/CV_project/api/models"
)

func HomeUsers(app *app.App, w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("Received request for /users")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	var users []models.User
	rows, err := app.DB.Query("SELECT id, firstname, lastname, email FROM users")
	if err != nil {
		app.Logger.Println("Error querying database: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email)
		if err != nil {
			app.Logger.Println("Error scanning row: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func Home(app *app.App, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	app.Logger.Println("Received request for /")

	rows, err := app.DB.Query("SELECT id, jobtitle, firstname, lastname, email, phone, address, city, country, postalcode, dateofbirth, nationality, summary, workexperience, education, skills, languages FROM users")
	if err != nil {
		app.Logger.Println("Error querying database: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Jobtitle, &user.Firstname, &user.Lastname, &user.Email, &user.Phone, &user.Address, &user.City, &user.Country, &user.Postalcode, &user.Dateofbirth, &user.Nationality, &user.Summary, &user.Workexperience, &user.Education, &user.Skills, &user.Languages)
		if err != nil {
			app.Logger.Println("Error scanning row: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
