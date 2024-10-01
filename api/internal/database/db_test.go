// * api/internal/database/db_test.go
package database

import (
	"database/sql"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// TestConnectToDatabases tests the ConnectToDatabases function
func TestConnectToDatabases(t *testing.T) {
    // Set up environment variables for test
    os.Setenv("MYSQL_USER", "testuser")
    os.Setenv("MYSQL_PASSWORD", "testpassword")
    os.Setenv("MYSQL_HOST", "localhost")
    os.Setenv("MYSQL_PORT", "3306")
    os.Setenv("MYSQL_DATABASE", "testdb")

    // Mock the database connection
    db, mock, err := sqlmock.New()
    assert.NoError(t, err)
    defer db.Close()

    // Override the sqlOpen function to return the mock database
    mockSQLOpen := func(driverName, dataSourceName string) (*sql.DB, error) {
        return db, nil
    }

    // Set up expectations for the mock database
    mock.ExpectPing()

    // Call the function to test
    conn, err := ConnectToDatabases(mockSQLOpen)
    assert.NoError(t, err)
    assert.NotNil(t, conn)

    // Ensure all expectations were met
    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}