package mock

import (
	"bytes"
	"time"
	_ "time/tzdata" // Blank import necessary for side effects

	"github.com/sirupsen/logrus"
)

// Logger is a mock logger for testing
type Logger struct {
	*logrus.Logger
}

// CustomFormatter formats logs without the log level
type CustomFormatter struct {
	TimestampFormat string
	ForceColors     bool
}

// Format formats the log entry
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b bytes.Buffer

	// Set color based on log level
	color := "\033[1;34;1m" // Default to blue
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		color = "\033[1;36;1m" // Cyan
	case logrus.InfoLevel:
		color = "\033[1;32;1m" // Green
	case logrus.WarnLevel:
		color = "\033[1;33;1m" // Yellow
	case logrus.ErrorLevel:
		color = "\033[1;31;1m" // Red
	case logrus.FatalLevel, logrus.PanicLevel:
		color = "\033[1;35;1m" // Magenta
	}
	if f.ForceColors {
		b.WriteString(color)
	}

	loc, err := time.LoadLocation("Europe/London")
	if err != nil {
		return nil, err
	}
	b.WriteString(entry.Time.In(loc).Format(f.TimestampFormat))
	b.WriteString(" ")
	b.WriteString(entry.Message)

	if f.ForceColors {
		b.WriteString("\033[0m")
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

// NewMockLogger creates a new instance of Logger with logrus.Logger initialized
func NewMockLogger() *Logger {
	logger := logrus.New()
	logger.SetFormatter(&CustomFormatter{
		TimestampFormat: "15:04:05", // Only show time
		ForceColors:     true,
	})
	return &Logger{Logger: logger}
}

// Print prints the log message
func (m *Logger) Print(v ...interface{}) {
	m.Logger.Print(v...)
}

// Printf prints the log message with formatting
func (m *Logger) Printf(format string, v ...interface{}) {
	m.Logger.Printf(format, v...)
}

// Println prints the log message with a newline
func (m *Logger) Println(v ...interface{}) {
	m.Logger.Println(v...)
}

// Fatal prints the log message and exits the program
func (m *Logger) Fatal(v ...interface{}) {
	m.Logger.Fatal(v...)
}

// Fatalf prints the log message with formatting and exits the program
func (m *Logger) Fatalf(format string, v ...interface{}) {
	m.Logger.Fatalf(format, v...)
}
