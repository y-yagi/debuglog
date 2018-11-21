package debuglog

import (
	"io"
	"log"
	"os"
)

// Logger is a logger for debug.
type Logger struct {
	logger *log.Logger
	debug  bool
}

// New creates a new Debug Logger.
func New(w io.Writer) *Logger {
	dl := &Logger{}
	dl.logger = log.New(w, "[DEBUG] ", 0)
	if len(os.Getenv("DEBUG")) != 0 {
		dl.debug = true
	}

	return dl
}

// Printf print log with format if DEBUG env specified.
func (l *Logger) Printf(format string, a ...interface{}) {
	if l.debug {
		l.logger.Printf(format, a...)
	}

	return
}

// Print print log if DEBUG env specified.
func (l *Logger) Print(a ...interface{}) {
	if l.debug {
		l.logger.Print(a...)
	}

	return
}
