// * handlers/templates_test.go
package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/CristyNel/CV_project/tree/main/api/mock"
	"github.com/DATA-DOG/go-sqlmock"
	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/stretchr/testify/assert"
)

func TestGenerateTemplate(t *testing.T) {
	app, mock := mock.SetupMockApp(t)
	defer app.DB.Close()

	// Mock the database query for template path
	mock.ExpectQuery("SELECT Path FROM template WHERE id = ?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"Path"}).AddRow("template_path.html"))

	// Mock the database query for user data
	mock.ExpectQuery(`SELECT \* FROM users WHERE id = \?`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "jobtitle", "firstname", "lastname", "email", "phone", "address", "city", "country", "postalcode", "dateofbirth", "nationality", "summary", "workexperience", "education", "skills", "languages"}).
			AddRow(1, "Developer", "John", "Doe", "john@example.com", "1234567890", "123 Street", "City", "Country", "12345", "1990-01-01", "Nationality", "Summary", "Work Experience", "Education", "Skills", "Languages"))

	// Mock file read
	htmlContent := "<html><body>{{Firstname}} {{Lastname}}</body></html>"
	err := os.WriteFile("template_path.html", []byte(htmlContent), 0644)
	if err != nil {
		t.Fatalf("failed to create template_path.html: %v", err)
	}
	defer os.Remove("template_path.html")

	// Ensure the directory exists
	populateTemplateDir := "./bff/templates/view/"
	absPath, err := filepath.Abs(populateTemplateDir)
	if err != nil {
		t.Fatalf("failed to get absolute path: %v", err)
	}
	fmt.Println("Absolute path:", absPath)
	err = os.MkdirAll(populateTemplateDir, os.ModePerm)
	if err != nil {
		t.Fatalf("failed to create directory: %v", err)
	}
	defer os.RemoveAll(populateTemplateDir)

	// Mock PDF generation
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		t.Fatalf("failed to create PDF generator: %v", err)
	}
	pdfg.AddPage(wkhtml.NewPageReader(bytes.NewReader([]byte(htmlContent))))
	if err := pdfg.Create(); err != nil {
		t.Fatalf("failed to create PDF: %v", err)
	}
	if err := pdfg.WriteFile("example.pdf"); err != nil {
		t.Fatalf("failed to write PDF file: %v", err)
	}
	defer os.Remove("example.pdf")

	req, err := http.NewRequest("GET", "/pdf?template=1&user=1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GenerateTemplate(app, w, r)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")
	assert.Contains(t, rr.Body.String(), "1, 1", "Expected response to contain template and user IDs")
}
