// main_test.go
package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"

	// "reflect"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var expectedUsers = []User{
	{
		ID:             1,
		Jobtitle:       "Developer",
		Firstname:      "John",
		Lastname:       "Doe",
		Email:          "john.doe@example.com",
		Phone:          "1234567890",
		Address:        "123 Street",
		City:           "City",
		Country:        "Country",
		Postalcode:     "12345",
		Dateofbirth:    "1990-01-01",
		Nationality:    "Nationality",
		Summary:        "Summary",
		Workexperience: "Work Experience",
		Education:      "Education",
		Skills:         "Skills",
		Languages:      "Languages",
	},
	{
		ID:             2,
		Jobtitle:       "Manager",
		Firstname:      "Jane",
		Lastname:       "Doe",
		Email:          "jane.doe@example.com",
		Phone:          "0987654321",
		Address:        "456 Avenue",
		City:           "City",
		Country:        "Country",
		Postalcode:     "67890",
		Dateofbirth:    "1985-01-01",
		Nationality:    "Nationality",
		Summary:        "Summary",
		Workexperience: "Work Experience",
		Education:      "Education",
		Skills:         "Skills",
		Languages:      "Languages",
	},
}

// mock handlers as methods on the App struct
func (app *App) mockHome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Home"))
}

func (app *App) mockHomeUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Users"))
}

// TestRouter tests the router
func TestRouter(t *testing.T) {
	app := &App{
		Logger: log.New(os.Stdout, "INFO: ", log.LstdFlags),
		DB:     &sql.DB{},
	}

	// Create a new router and register the handlers
	r := mux.NewRouter()
	r.HandleFunc("/", app.mockHome).Methods("GET")
	r.HandleFunc("/users", app.mockHomeUsers).Methods("GET")

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "Users"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// * verify login
func TestVerifyLogin(t *testing.T) {
	app := &App{
		Logger: log.New(os.Stdout, "INFO: ", log.LstdFlags),
		DB:     &sql.DB{},
	}

	var err error
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	// override db with mockDB
	originalDB := app.DB
	app.DB = mockDB
	defer func() { app.DB = originalDB }()

	// generate hashed password
	password := "password123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when hashing a password", err)
	}

	tests := []struct {
		name     string
		username string
		password string
		mock     func()
		want     bool
	}{
		{
			name:     "Valid login",
			username: "user1",
			password: "password123",
			mock: func() {
				mock.ExpectQuery("SELECT password FROM userlogin WHERE username = ?").
					WithArgs("user1").
					WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow(string(hashedPassword)))
			},
			want: true,
		},
		{
			name:     "Invalid login",
			username: "user1",
			password: "wrongpassword",
			mock: func() {
				mock.ExpectQuery("SELECT password FROM userlogin WHERE username = ?").
					WithArgs("user1").
					WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow(string(hashedPassword)))
			},
			want: false,
		},
		{
			name:     "Database error",
			username: "user1",
			password: "password123",
			mock: func() {
				mock.ExpectQuery("SELECT password FROM userlogin WHERE username = ?").
					WithArgs("user1").
					WillReturnError(errors.New("db error"))
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got := app.VerifyLogin(tt.username, tt.password)
			if got != tt.want {
				t.Errorf("verifyLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}

// set session
func TestSetSession(t *testing.T) {
	app := &App{
		Logger: log.New(os.Stdout, "INFO: ", log.LstdFlags),
		DB:     &sql.DB{},
	}

	userName := "testUser"
	recorder := httptest.NewRecorder()

	app.SetSession(userName, recorder)

	result := recorder.Result()
	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", result.Status)
	}

	// Check if the session cookie is set
	cookies := result.Cookies()
	if len(cookies) == 0 {
		t.Errorf("expected session cookie to be set; got none")
	}

	var sessionCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "session" {
			sessionCookie = cookie
			break
		}
	}

	if sessionCookie == nil {
		t.Errorf("expected session cookie to be set; got none")
	} else {
		if sessionCookie.Value == "" {
			t.Errorf("expected session cookie to have a value; got empty")
		}
	}
}

// logout handler
func TestLogoutHandler_ClearsSession(t *testing.T) {
	app := &App{
		Logger: log.New(os.Stdout, "INFO: ", log.LstdFlags),
		DB:     &sql.DB{},
	}

	req, err := http.NewRequest("GET", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.LogoutHandler)

	handler.ServeHTTP(rr, req)

	cookies := rr.Result().Cookies()
	if len(cookies) == 0 {
		t.Errorf("Expected session cookie to be cleared, got none")
	}

	var sessionCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "session" {
			sessionCookie = cookie
			break
		}
	}

	if sessionCookie == nil {
		t.Errorf("Expected session cookie to be cleared, got none")
	} else {
		if sessionCookie.Value != "" || sessionCookie.MaxAge != -1 {
			t.Errorf("Expected session cookie to be cleared, got %v", sessionCookie)
		}
	}
}

func TestLogoutHandler_StatusNoContent(t *testing.T) {
	app := &App{
		Logger: log.New(os.Stdout, "INFO: ", log.LstdFlags),
		DB:     &sql.DB{},
	}

	req, err := http.NewRequest("GET", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.LogoutHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("Expected status code %v, got %v", http.StatusNoContent, status)
	}
}

func TestLogoutHandler_ResponseMessage(t *testing.T) {
	app := &App{
		Logger: log.New(os.Stdout, "INFO: ", log.LstdFlags),
		DB:     &sql.DB{},
	}

	req, err := http.NewRequest("GET", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.LogoutHandler)

	handler.ServeHTTP(rr, req)

	expected := "User logout successfully"
	if rr.Body.String() != expected {
		t.Errorf("Expected response body %v, got %v", expected, rr.Body.String())
	}
}

// signup handler
func TestSignupHandler_ParseFormError(t *testing.T) {
	app := &App{
		Logger: log.New(os.Stdout, "INFO: ", log.LstdFlags),
		DB:     &sql.DB{},
	}

	req, err := http.NewRequest("POST", "/signup", strings.NewReader("invalid=form"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.SignupHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, status)
	}
}

func TestSignupHandler_EmptyEmailOrPassword(t *testing.T) {
	app := &App{
		Logger: log.New(os.Stdout, "INFO: ", log.LstdFlags),
		DB:     &sql.DB{},
	}

	req, err := http.NewRequest("POST", "/signup", strings.NewReader("email=&password="))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.SignupHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, status)
	}
}

func TestHomeHandler(t *testing.T) {
	// Create a mock database
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	// Initialize the App struct with the mock database and a logger
	app := &App{
		DB:     mockDB,
		Logger: log.New(os.Stdout, "INFO: ", log.LstdFlags),
	}

	// Mock the database query
	rows := sqlmock.NewRows([]string{"id", "jobtitle", "firstname", "lastname", "email", "phone", "address", "city", "country", "postalcode", "dateofbirth", "nationality", "summary", "workexperience", "education", "skills", "languages"}).
		AddRow(1, "Developer", "John", "Doe", "john.doe@example.com", "1234567890", "123 Street", "City", "Country", "12345", "1990-01-01", "Nationality", "Summary", "Work Experience", "Education", "Skills", "Languages").
		AddRow(2, "Manager", "Jane", "Doe", "jane.doe@example.com", "0987654321", "456 Avenue", "City", "Country", "67890", "1985-01-01", "Nationality", "Summary", "Work Experience", "Education", "Skills", "Languages")

	mock.ExpectQuery("SELECT id, jobtitle, firstname, lastname, email, phone, address, city, country, postalcode, dateofbirth, nationality, summary, workexperience, education, skills, languages FROM users").WillReturnRows(rows)

	req, err := http.NewRequest("GET", "/home", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.Home)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var gotUsers []User
	if err := json.NewDecoder(rr.Body).Decode(&gotUsers); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if !reflect.DeepEqual(gotUsers, expectedUsers) {
		t.Errorf("handler returned unexpected body: got %v want %v", gotUsers, expectedUsers)
	}
}

func TestHomeHandler_NonGETMethod(t *testing.T) {
	app := &App{
		Logger: log.New(os.Stdout, "INFO: ", log.LstdFlags),
		DB:     &sql.DB{}, // Mock or initialize your DB connection here
	}
	req, err := http.NewRequest("POST", "/home", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.Home)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}

	expected := "Method Not Allowed"
	actual := strings.TrimSpace(rr.Body.String())
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
