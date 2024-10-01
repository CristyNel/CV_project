// * api/mock/mock_services.go
package mock

import (
	"net/http"
)

// MockServiceHandler is a mock handler for external services
func MockServiceHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Mock service response"}`))
}
