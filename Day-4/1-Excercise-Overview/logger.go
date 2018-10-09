package main

import (
	"fmt"
	"sync"
	"io"
	"bytes"
	"os"
)

// Define levels
type Level int
const (
	LevelDebug Level = iota
	LevelInfo
	LevelError
)

const defaultLogLevel = LevelInfo

// Type Logger represents a logging object that handles log entries based on the current log level
type Logger struct {
	mu         sync.Mutex // for serialization
	prefix     string     // prefix to write at beginning of each log entry
	Level      Level
	w          io.Writer    // writer for output
	buf        bytes.Buffer // internal buffer
}

// New creates a new Logger.
func New(w io.Writer, prefix string) *Logger {
	return &Logger{w: w, prefix: prefix, Level: defaultLogLevel }
}

// Console creates a new Logger that outputs to Stderr.
var Console = New(os.Stderr, "")

func (l *Logger) Debug(v ...interface{}) {
	if LevelDebug < l.Level {
		return
	}
	l.WriteEntry(LevelDebug, fmt.Sprintln(v...))
}

func (l *Logger) Info(v ...interface{}) {
	if LevelInfo < l.Level {
		return
	}
	l.WriteEntry(LevelInfo, fmt.Sprintln(v...))
}

func (l *Logger) Error(v ...interface{}) {
	if LevelError < l.Level {
		return
	}
	l.WriteEntry(LevelError, fmt.Sprintln(v...))
}

// SetLevel sets the output level for the logger.
func (l *Logger) SetLevel(lvl Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Level = lvl
}

// GetLevel gets the output level of the logger.
func (l *Logger) GetLevel() Level {
	return l.Level
}

// WriteEntry writes the msg of the specified level to the underling writer
func (l *Logger) WriteEntry(lvl Level, msg string) error {
	l.w.Write([]byte(msg))
	return nil
}

func main()  {
	Console.Info("Hello")
	Console.Debug("Hello", "Debugger")
	Console.Error("Error")

	Console.SetLevel(LevelError)

	Console.Info("Hello")
	Console.Debug("Hello", "Debugger")
	Console.Error("Error")
}

