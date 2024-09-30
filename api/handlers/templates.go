// handlers/templates.go
package handlers

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/damarisnicolae/CV_project/api/internal/app"
	"github.com/damarisnicolae/CV_project/api/internal/utils"
	"github.com/damarisnicolae/CV_project/api/models"
)

func GenerateTemplate(app *app.App, w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("Received request for /pdf")

	query := r.URL.Query()
	template_id := query["template"]
	user_id := query["user"]

	iduser_int, err := strconv.Atoi(user_id[0])
	if err != nil {
		log.Printf("An error occurred: %v", err)
	}

	idtemplate_int, err := strconv.Atoi(template_id[0])
	if err != nil {
		log.Printf("An error occurred: %v", err)
	}

	var user models.User
	var template utils.Template

	// Get the path of the template
	row1 := app.DB.QueryRow("SELECT Path FROM template WHERE id = ?", idtemplate_int)

	if err := row1.Scan(&template.Path); err != nil {
		if err == sql.ErrNoRows {
			app.Logger.Println("Error scanning row: ", err)
			http.NotFound(w, r)
			return
		}
		http.Error(w, fmt.Sprintf("Error fetching user data: %v", err), http.StatusInternalServerError)
		return
	}

	row := app.DB.QueryRow("SELECT * FROM users WHERE id = ?", iduser_int)

	if err := row.Scan(&user.ID, &user.Jobtitle, &user.Firstname, &user.Lastname, &user.Email, &user.Phone, &user.Address, &user.City, &user.Country, &user.Postalcode, &user.Dateofbirth, &user.Nationality, &user.Summary, &user.Workexperience, &user.Education, &user.Skills, &user.Languages); err != nil {
		if err == sql.ErrNoRows {
			app.Logger.Println("Error scanning row: ", err)
			http.NotFound(w, r)
			return
		}
		http.Error(w, fmt.Sprintf("Error fetching user data: %v", err), http.StatusInternalServerError)
		return
	}

	htmlContent, err := os.ReadFile(template.Path)
	if err != nil {
		panic(err)
	}

	htmlString := string(htmlContent)

	htmlString = strings.ReplaceAll(htmlString, "{{Firstname}}", user.Firstname)
	htmlString = strings.ReplaceAll(htmlString, "{{Lastname}}", user.Lastname)
	htmlString = strings.ReplaceAll(htmlString, "{{Jobtitle}}", user.Jobtitle)
	htmlString = strings.ReplaceAll(htmlString, "{{Email}}", user.Email)
	htmlString = strings.ReplaceAll(htmlString, "{{Phone}}", user.Phone)
	htmlString = strings.ReplaceAll(htmlString, "{{Address}}", user.Address)
	htmlString = strings.ReplaceAll(htmlString, "{{City}}", user.City)
	htmlString = strings.ReplaceAll(htmlString, "{{Country}}", user.Country)
	htmlString = strings.ReplaceAll(htmlString, "{{Postalcode}}", user.Postalcode)
	htmlString = strings.ReplaceAll(htmlString, "{{Dateofbirth}}", user.Dateofbirth)
	htmlString = strings.ReplaceAll(htmlString, "{{Nationality}}", user.Nationality)
	htmlString = strings.ReplaceAll(htmlString, "{{Summary}}", user.Summary)
	htmlString = strings.ReplaceAll(htmlString, "{{Workexperience}}", user.Workexperience)
	htmlString = strings.ReplaceAll(htmlString, "{{Education}}", user.Education)
	htmlString = strings.ReplaceAll(htmlString, "{{Skills}}", user.Skills)
	htmlString = strings.ReplaceAll(htmlString, "{{Languages}}", user.Languages)

	// Write
	err = os.WriteFile("../bff/templates/populate_template.html", []byte(htmlString), 0644)
	if err != nil {
		panic(err)
	}
	// Read
	populateHtml, err := os.ReadFile("../bff/templates/populate_template.html")
	if err != nil {
		log.Fatal(err)
	}
	// Create PDF
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		return
	}
	// Add HTML page
	pdfg.AddPage(wkhtml.NewPageReader(bytes.NewReader(populateHtml)))
	// Create the PDF document in memory
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}
	// Write the PDF document to a file
	err = pdfg.WriteFile("./example.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// Respond with template and user IDs
	fmt.Fprintf(w, "%s, %s", template_id, user_id)
}
