// * api/models/user_test.go
package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserJSONMarshaling(t *testing.T) {
	user := User{
		ID:             1,
		Jobtitle:       "Developer",
		Firstname:      "John",
		Lastname:       "Doe",
		Email:          "john.doe@example.com",
		Phone:          "1234567890",
		Address:        "123 Main St",
		City:           "Anytown",
		Country:        "USA",
		Postalcode:     "12345",
		Dateofbirth:    "1990-01-01",
		Nationality:    "American",
		Summary:        "Experienced developer",
		Workexperience: "5 years",
		Education:      "Bachelor's Degree",
		Skills:         "Go, Java, Python",
		Languages:      "English, Spanish",
	}

	// Test JSON marshaling
	jsonData, err := json.Marshal(user)
	assert.NoError(t, err)

	expectedJSON := `{"id":1,"jobtitle":"Developer","firstname":"John","lastname":"Doe","email":"john.doe@example.com","phone":"1234567890","address":"123 Main St","city":"Anytown","country":"USA","postalcode":"12345","dateofbirth":"1990-01-01","nationality":"American","summary":"Experienced developer","workexperience":"5 years","education":"Bachelor's Degree","skills":"Go, Java, Python","languages":"English, Spanish"}`
	assert.JSONEq(t, expectedJSON, string(jsonData))

	// Test JSON unmarshaling
	var unmarshaledUser User
	err = json.Unmarshal(jsonData, &unmarshaledUser)
	assert.NoError(t, err)
	assert.Equal(t, user, unmarshaledUser)
}
