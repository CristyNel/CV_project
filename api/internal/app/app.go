// * CV_project/api/internal/app/app.go
package app

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

// LoggerInterface defines the methods that a logger should implement
type LoggerInterface interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

// App holds the application-wide dependencies
type App struct {
	DB     *sql.DB
	Logger *logrus.Logger
}
