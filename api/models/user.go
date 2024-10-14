// * CV_project/api/models/user.go

package models

// User represents a user in the system
type User struct {
	ID             int    `json:"id"`
	Jobtitle       string `json:"jobtitle"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
	City           string `json:"city"`
	Country        string `json:"country"`
	Postalcode     string `json:"postalcode"`
	Dateofbirth    string `json:"dateofbirth"`
	Nationality    string `json:"nationality"`
	Summary        string `json:"summary"`
	Workexperience string `json:"workexperience"`
	Education      string `json:"education"`
	Skills         string `json:"skills"`
	Languages      string `json:"languages"`
}
