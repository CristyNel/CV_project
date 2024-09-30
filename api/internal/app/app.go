// CV_project/api/internal/app/app.go

package app

import (
	"database/sql"
	"log"
)

// App holds the application-wide dependencies
type App struct {
	DB     *sql.DB
	Logger *log.Logger
}
