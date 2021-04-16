package debuglog

import (
	"io"
	"log"
	"os"
)

// Logger is a logger for debug.
type Logger struct {
	logger *log.Logger
	envkey string
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

// EnvKey is a flag for specifying the environment key.
func EnvKey(k string) Option {
	return func(l *Logger) error {
		l.envkey = k
		return nil
	}
}

// New creates a new Debug Logger.
func New(w io.Writer, options ...Option) *Logger {
	dl := &Logger{envkey: "DEBUG"}
	dl.logger = log.New(w, "[DEBUG] ", 0)

	for _, option := range options {
		option(dl)
	}

	if len(os.Getenv(dl.envkey)) != 0 {
		dl.debug = true
	}

	return dl
}

// Printf print log with format if DEBUG env specified.
func (l *Logger) Printf(format string, a ...interface{}) {
	if l.debug {
		l.logger.Printf(format, a...)
	}
}

// Print print log if DEBUG env specified.
func (l *Logger) Print(a ...interface{}) {
	if l.debug {
		l.logger.Print(a...)
	}
}
