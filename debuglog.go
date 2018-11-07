package debuglog

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// Logger is a logger for debug.
type Logger struct {
	out    io.Writer
	debug  bool
	prefix string
	mu     sync.Mutex
}

// New creates a new Debug Logger.
func New(w io.Writer) *Logger {
	dl := &Logger{out: w, prefix: "[DEBUG] "}

	if len(os.Getenv("DEBUG")) != 0 {
		dl.debug = true
	}

	return dl
}

// Printf print log with format if DEBUG env specified.
func (l *Logger) Printf(format string, a ...interface{}) (n int, err error) {
	if l.debug {
		l.mu.Lock()
		defer l.mu.Unlock()
		return fmt.Fprintf(l.out, l.prefix+format, a...)
	}

	return 0, nil
}

// Print print log if DEBUG env specified.
func (l *Logger) Print(a ...interface{}) (n int, err error) {
	if l.debug {
		a = append([]interface{}{l.prefix}, a...)
		l.mu.Lock()
		defer l.mu.Unlock()
		return fmt.Fprint(l.out, a...)
	}

	return 0, nil
}
