// CV_project/api/cmd/main.go

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/damarisnicolae/CV_project/api/internal/app"
	"github.com/damarisnicolae/CV_project/api/internal/database"
	"github.com/damarisnicolae/CV_project/api/routes"
)

func main() {
	var err error

	Db, err := database.ConnectToDatabases()
	if err != nil {
		log.Printf("\033[1;31;1m * Failed to connect to the database: %v\033[0m", err)
		return
	}
	defer Db.Close()

	if Db == nil {
		log.Fatalf("\033[1;31;1m * Failed to initialize the database connection.\033[0m")
	}

	app := &app.App{
		DB:     Db,
		Logger: log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}

	r := routes.InitializeRouter(app)

	app.Logger.Println("\n\033[1;37;1m * Starting the HTTP server on port: ➮\033[1;94;1m 8080\033[0m")
	if err := http.ListenAndServe(":8080", r); err != nil {
		app.Logger.Fatalf("\n * Failed to start HTTP server: %s\n", err)
	}
}
