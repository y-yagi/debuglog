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

// Option is a option for logger.
type Option func(*Logger) error

// Flag is a flag for logger. Can specify same flag of log package.
func Flag(f int) Option {
	return func(l *Logger) error {
		l.logger.SetFlags(f)
		return nil
	}
}

// New creates a new Debug Logger.
func New(w io.Writer, options ...Option) *Logger {
	dl := &Logger{}
	dl.logger = log.New(w, "[DEBUG] ", 0)
	if len(os.Getenv("DEBUG")) != 0 {
		dl.debug = true
	}

	for _, option := range options {
		option(dl)
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
