// * api/mock/mock_session_test.go
package mock

import (
    "net/http"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestNewMockSession(t *testing.T) {
    rr := NewMockSession()

    // Simulate setting a session value
    http.SetCookie(rr, &http.Cookie{Name: "session_id", Value: "mock_session_value"})

    // Check if the session value is set correctly
    cookies := rr.Result().Cookies()
    assert.NotEmpty(t, cookies, "Expected cookies to be set")
    assert.Equal(t, "mock_session_value", cookies[0].Value, "Expected session value to be 'mock_session_value'")
}