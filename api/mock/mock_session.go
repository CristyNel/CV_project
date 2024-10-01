// * api/mock/mock_session.go
package mock

import (
	"net/http/httptest"
)

// NewMockSession creates a new mock session for testing
func NewMockSession() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}
