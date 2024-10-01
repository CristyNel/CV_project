// * CV_project/api/internal/app/app_test.go
package app

import (
	"database/sql"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestAppInitialization(t *testing.T) {
	// Create mock instances of sql.DB and logrus.Logger
	mockDB := &sql.DB{}
	mockLogger := logrus.New()

	// Initialize the App struct with the mock instances
	app := &App{
		DB:     mockDB,
		Logger: mockLogger,
	}

	// Verify that the App struct holds the correct instances
	assert.Equal(t, mockDB, app.DB)
	assert.Equal(t, mockLogger, app.Logger)
}
