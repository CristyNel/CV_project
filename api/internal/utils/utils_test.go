// * api/internal/utils/utils_test.go
package utils

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/damarisnicolae/CV_project/api/internal/app"
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
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	app := &app.App{DB: db}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	mock.ExpectQuery("SELECT password FROM userlogin WHERE username = ?").
		WithArgs("user").
		WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow(string(hashedPassword)))

	assert.True(t, VerifyLogin(app, "user", "password"))
	assert.False(t, VerifyLogin(app, "user", "wrongpassword"))

	mock.ExpectQuery("SELECT password FROM userlogin WHERE username = ?").
		WithArgs("nonexistent").
		WillReturnError(sql.ErrNoRows)

	assert.False(t, VerifyLogin(app, "nonexistent", "password"))
}

// TestSetSession tests the SetSession function
func TestSetSession(t *testing.T) {
	app := &app.App{}
	w := httptest.NewRecorder()
	SetSession(app, "testuser", w)

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
