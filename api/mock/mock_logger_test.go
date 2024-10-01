// * CV_project/api/mock/mock_logger_test.go
package mock

import (
	"bytes"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewMockLogger(t *testing.T) {
	logger := NewMockLogger()
	assert.NotNil(t, logger)
	assert.IsType(t, &Logger{}, logger)
}

func TestLogger_Print(t *testing.T) {
	var buf bytes.Buffer
	logger := logrus.New()
	logger.Out = &buf

	logger.Print("test message")
	assert.Contains(t, buf.String(), "test message")
}

func TestLogger_Printf(t *testing.T) {
	var buf bytes.Buffer
	logger := logrus.New()
	logger.Out = &buf

	logger.Printf("test %s", "message")
	assert.Contains(t, buf.String(), "test message")
}

func TestLogger_Println(t *testing.T) {
	var buf bytes.Buffer
	logger := logrus.New()
	logger.Out = &buf

	logger.Println("test message")
	assert.Contains(t, buf.String(), "test message")
}
