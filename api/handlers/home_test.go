// * api/handlers/home_test.go
package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CristyNel/CV_project/tree/main/api/mock"
	"github.com/CristyNel/CV_project/tree/main/api/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	app, mock := mock.SetupMockApp(t)
	defer app.DB.Close()

	// Mock the database query
	rows := sqlmock.NewRows([]string{"id", "jobtitle", "firstname", "lastname", "email", "phone", "address", "city", "country", "postalcode", "dateofbirth", "nationality", "summary", "workexperience", "education", "skills", "languages"}).
		AddRow(1, "Developer", "John", "Doe", "john@example.com", "1234567890", "123 Street", "City", "Country", "12345", "1990-01-01", "Nationality", "Summary", "Work Experience", "Education", "Skills", "Languages").
		AddRow(2, "Designer", "Jane", "Doe", "jane@example.com", "0987654321", "456 Avenue", "City", "Country", "67890", "1992-02-02", "Nationality", "Summary", "Work Experience", "Education", "Skills", "Languages")

	mock.ExpectQuery("SELECT id, jobtitle, firstname, lastname, email, phone, address, city, country, postalcode, dateofbirth, nationality, summary, workexperience, education, skills, languages FROM users").
		WillReturnRows(rows)

	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Home(app, w, r)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")

	var users []models.User
	err = json.NewDecoder(rr.Body).Decode(&users)
	assert.NoError(t, err)

	assert.Len(t, users, 2, "Expected 2 users")
	assert.Equal(t, "John", users[0].Firstname, "Expected first user to be John")
	assert.Equal(t, "Jane", users[1].Firstname, "Expected second user to be Jane")
}

func TestHomeUsersHandler(t *testing.T) {
	app, mock := mock.SetupMockApp(t)
	defer app.DB.Close()

	// Mock the database query
	rows := sqlmock.NewRows([]string{"id", "firstname", "lastname", "email"}).
		AddRow(1, "John", "Doe", "john@example.com").
		AddRow(2, "Jane", "Doe", "jane@example.com")

	mock.ExpectQuery("SELECT id, firstname, lastname, email FROM users").
		WillReturnRows(rows)

	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		HomeUsers(app, w, r)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")

	var users []models.User
	err = json.NewDecoder(rr.Body).Decode(&users)
	assert.NoError(t, err)

	assert.Len(t, users, 2, "Expected 2 users")
	assert.Equal(t, "John", users[0].Firstname, "Expected first user to be John")
	assert.Equal(t, "Jane", users[1].Firstname, "Expected second user to be Jane")
}
