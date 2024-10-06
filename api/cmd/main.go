// CV_project/api/cmd/main.go

package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CristyNel/CV_project/tree/main/api/internal/app"
	"github.com/CristyNel/CV_project/tree/main/api/internal/database"
	"github.com/CristyNel/CV_project/tree/main/api/mock"
	"github.com/CristyNel/CV_project/tree/main/api/routes"
)

func main() {
	var err error

	Db, err := database.ConnectToDatabases(sql.Open)
	if err != nil {
		log.Printf("\033[1;31;1m * * * 🚨 Failed to connect to the database: %v\033[0m", err)
		return
	}
	defer Db.Close()

	if Db == nil {
		log.Fatalf("\033[1;31;1m * * * 🚨 Failed to initialize the database connection.\033[0m")
	}

	// Use the mock logger
	logger := mock.NewMockLogger()
	app := &app.App{
		DB:     Db,
		Logger: logger,
	}

	r := routes.InitializeRouter(app)

	app.Logger.Printf("\n\033[1;37;1m * * * 🛫 Starting the HTTP server on port: ➮\033[1;94;1m 8080\033[0m")
	if err := http.ListenAndServe(":8080", r); err != nil {
		app.Logger.Fatalf("\n * Failed to start HTTP server: %s\n", err)
	}
}
