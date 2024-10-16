package mock

import (
	"database/sql"
	"testing"

	"github.com/CristyNel/CV_project/tree/main/api/internal/app"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"
)

// NewMockDB creates a new mock database connection for testing
func NewMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating sqlmock: %v", err)
	}
	return db, mock
}

// SetupMockApp creates a mock application environment for testing
func SetupMockApp(t *testing.T) (*app.App, sqlmock.Sqlmock) {
	db, mock := NewMockDB(t)
	logger := logrus.New()
	return &app.App{
		DB:     db,
		Logger: logger,
	}, mock
}
