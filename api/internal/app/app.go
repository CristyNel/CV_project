// * CV_project/api/internal/app/app.go
package app

import (
	"database/sql"
)

// LoggerInterface defines the methods that a logger should implement
type Logger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}

// App holds the application-wide dependencies
type App struct {
	DB     *sql.DB
	Logger Logger
}
