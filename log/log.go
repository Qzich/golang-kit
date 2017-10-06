package log

import (
	"fmt"
	"io"
	"sync"
)

//
// Notifications severities
//
const (
	SeverityError = "ERROR"
	SeverityInfo  = "INFO"
	SeverityDebug = "DEBUG"
)

//
// Logger interface is the interface for the Log object.
// ITODO: Think on two logging streams: regular and stacktrace one
//
// noinspection GoNameStartsWithPackageName
type Logger interface {
	Error(format string, args ...interface{})
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
}

//
// Log object constructor
//
// ITODO: reuse system logger class
// Log is deliberately named is this way because it is self-explanatory.
//
// noinspection GoNameStartsWithPackageName
type Log struct {
	writer   io.Writer
	severity string
	mu       sync.RWMutex
}

//
// New returns an instance of a logger.
//
func New(writer io.Writer, severity string) *Log {
	return &Log{
		writer:   writer,
		severity: severity,
	}
}

//
// Error logs an error.
//
func (l *Log) Error(format string, args ...interface{}) {
	l.write(SeverityError, format, args...)
}

//
// Errorf logs a formatted error info.
//
func (l *Log) Errorf(format string, args ...interface{}) {

	l.write(SeverityError, fmt.Sprintf(format, args...))
}

//
// Debug logs a formatted debug info.
//
func (l *Log) Debug(format string, args ...interface{}) {

	l.write(SeverityDebug, fmt.Sprintf(format, args...))
}

//
// Info logs a formatted info.
//
func (l *Log) Info(format string, args ...interface{}) {

	l.write(SeverityInfo, format, args...)
}

//
// write Writes the string to the output.
// TODO: Add datetime and function name
//
func (l *Log) write(severity string, format string, args ...interface{}) {
	l.mu.RLock()
	{
		m := fmt.Sprintf(format, args...)
		message := fmt.Sprintf("[%s] %s\n", severity, m)
		l.writer.Write([]byte(message))
	}
	l.mu.RUnlock()
}
