package app

import (
	"database/sql"

	"github.com/gorilla/sessions"
)

// Logger is an interface that defines the logging methods
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
	Store  *sessions.CookieStore
}
