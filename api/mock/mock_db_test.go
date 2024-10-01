// * CV_project/api/mock/mock_db_test.go
package mock

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNewMockDB(t *testing.T) {
	db, mock := NewMockDB(t)
	defer db.Close()

	// Ensure the mock expectations are met
	mock.ExpectQuery("SELECT 1").WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))

	// Execute a query to test the mock
	rows, err := db.Query("SELECT 1")
	assert.NoError(t, err)
	defer rows.Close()

	// Check if the query returned the expected result
	assert.True(t, rows.Next())
}

func TestSetupMockApp(t *testing.T) {
	app, mock := SetupMockApp(t)
	defer app.DB.Close()

	// Ensure the mock expectations are met
	mock.ExpectQuery("SELECT 1").WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))

	// Execute a query to test the mock
	rows, err := app.DB.Query("SELECT 1")
	assert.NoError(t, err)
	defer rows.Close()

	// Check if the query returned the expected result
	assert.True(t, rows.Next())

	// Check if the logger is properly initialized
	assert.NotNil(t, app.Logger)
}
