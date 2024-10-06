// * api/internal/utils/utils_test.go
package utils

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/CristyNel/CV_project/tree/main/api/internal/app"
	"github.com/CristyNel/CV_project/tree/main/api/mock"
	"github.com/gorilla/securecookie"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

// TestGetEnv tests the GetEnv function
func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_KEY", "test_value")
	app := &app.App{}
	value := GetEnv(app, "TEST_KEY", "default_value")
	assert.Equal(t, "test_value", value)

	value = GetEnv(app, "NON_EXISTENT_KEY", "default_value")
	assert.Equal(t, "default_value", value)
}

// TestVerifyLogin tests the VerifyLogin function
func TestVerifyLogin(t *testing.T) {
	// Initialize the app with a mock database and logger
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Create a mock logger
	logger := log.New(os.Stdout, "test: ", log.LstdFlags)

	// Initialize the app
	app := &app.App{
		DB:     db,
		Logger: logger,
	}

	// Create a mock user table and insert a test user
	_, err = db.Exec("CREATE TABLE user_session (username TEXT, password TEXT)")
	if err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}

	// Insert a test user with a hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("testpassword"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}
	_, err = db.Exec("INSERT INTO user_session (username, password) VALUES (?, ?)", "testuser", string(hashedPassword))
	if err != nil {
		t.Fatalf("Failed to insert test user: %v", err)
	}

	// Run the test
	if !VerifyLogin(app, "testuser", "testpassword") {
		t.Error("Expected login to succeed for valid credentials")
	}

	if VerifyLogin(app, "testuser", "wrongpassword") {
		t.Error("Expected login to fail for invalid credentials")
	}

	if VerifyLogin(app, "nonexistentuser", "testpassword") {
		t.Error("Expected login to fail for nonexistent user")
	}
}

// TestSetSession tests the SetSession function
func TestSetSession(t *testing.T) {
	// Initialize CookieHandler
	CookieHandler = securecookie.New(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32),
	)

	// Create a mock response recorder
	w := httptest.NewRecorder()

	// Create a mock app with the mock logger
	mockApp := &app.App{
		Logger: mock.NewMockLogger(),
	}

	// Call SetSession
	SetSession(mockApp, "testuser", w)

	// Check if the cookie is correctly set
	cookie := w.Result().Cookies()[0]
	assert.Equal(t, "session", cookie.Name)
	assert.NotEmpty(t, cookie.Value)
}

// TestErrorResponse tests the ErrorResponse function
func TestErrorResponse(t *testing.T) {
	w := httptest.NewRecorder()
	ErrorResponse(w, http.StatusBadRequest, "Bad Request")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error": "Bad Request"}`, w.Body.String())
}

// TestJSONResponse tests the JSONResponse function
func TestJSONResponse(t *testing.T) {
	w := httptest.NewRecorder()
	payload := map[string]string{"message": "success"}
	JSONResponse(w, http.StatusOK, payload)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "success"}`, w.Body.String())
}
