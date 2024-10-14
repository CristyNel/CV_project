// * api/mock/mock_services_test.go
package mock

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockServiceHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/mock-service", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ServiceHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")
	expected := `{"message":"Mock service response"}`
	assert.JSONEq(t, expected, rr.Body.String(), "Expected JSON response")
}
