// * CV_project/api/mock/mock_logger.go
package mock

import "github.com/sirupsen/logrus"

// Logger is a mock logger for testing
type Logger struct {
    *logrus.Logger
}

// NewMockLogger creates a new instance of Logger with logrus.Logger initialized
func NewMockLogger() *Logger {
    return &Logger{
        Logger: logrus.New(),
    }
}

func (m *Logger) Print(v ...interface{}) {
    m.Logger.Print(v...)
}

func (m *Logger) Printf(format string, v ...interface{}) {
    m.Logger.Printf(format, v...)
}

func (m *Logger) Println(v ...interface{}) {
    m.Logger.Println(v...)
}