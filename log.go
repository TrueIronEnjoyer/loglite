package loglite

import (
    "fmt"
    "os"
    "time"
)

// LogLevel defines the available logging levels.
type LogLevel int

const (
    INFO LogLevel = iota
    WARN
    ERROR
)

// Logger represents a simple logger.
type Logger struct {
    level  LogLevel
    output *os.File
}

// NewLogger creates a new Logger instance.
func NewLogger(level LogLevel, outputPath string) (*Logger, error) {
    var output *os.File
    var err error
    if outputPath != "" {
        output, err = os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
        if err != nil {
            return nil, err
        }
    } else {
        output = os.Stdout
    }
    return &Logger{level: level, output: output}, nil
}

// logMessage logs a message with the given level.
func (l *Logger) logMessage(level LogLevel, msg string) {
    if level < l.level {
        return
    }
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    levelStr := [...]string{"INFO", "WARN", "ERROR"}[level]
    logMsg := fmt.Sprintf("[%s] [%s] %s\n", timestamp, levelStr, msg)
    l.output.WriteString(logMsg)
}

// Info logs an informational message.
func (l *Logger) Info(msg string) {
    l.logMessage(INFO, msg)
}

// Warn logs a warning message.
func (l *Logger) Warn(msg string) {
    l.logMessage(WARN, msg)
}

// Error logs an error message.
func (l *Logger) Error(msg string) {
    l.logMessage(ERROR, msg)
}

// Close closes the logger's output file if it was opened.
func (l *Logger) Close() {
    if l.output != os.Stdout {
        l.output.Close()
    }
}
