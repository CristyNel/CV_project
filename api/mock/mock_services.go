package mock

import (
	"net/http"
)

// ServiceHandler is a mock handler for external services
func ServiceHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Mock service response"}`))
}
