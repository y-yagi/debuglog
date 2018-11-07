package debuglog

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// DebugLogger is a logger for debug.
type DebugLogger struct {
	out    io.Writer
	debug  bool
	prefix string
	mu     sync.Mutex
}

// New creates a new Debug Logger.
func New(w io.Writer) *DebugLogger {
	dl := &DebugLogger{out: w, prefix: "[DEBUG] "}

	if len(os.Getenv("DEBUG")) != 0 {
		dl.debug = true
	}

	return dl
}

// Printf print log with format if DEBUG env specified.
func (dl *DebugLogger) Printf(format string, a ...interface{}) (n int, err error) {
	if dl.debug {
		dl.mu.Lock()
		defer dl.mu.Unlock()
		return fmt.Fprintf(dl.out, dl.prefix+format, a...)
	}

	return 0, nil
}

// Print print log if DEBUG env specified.
func (dl *DebugLogger) Print(a ...interface{}) (n int, err error) {
	if dl.debug {
		a = append([]interface{}{dl.prefix}, a...)
		dl.mu.Lock()
		defer dl.mu.Unlock()
		return fmt.Fprint(dl.out, a...)
	}

	return 0, nil
}
