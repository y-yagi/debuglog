package debuglog

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func TestDebugEnvSpecified(t *testing.T) {
	os.Setenv("DEBUG", "1")

	out := new(bytes.Buffer)
	dl := New(out)

	dl.Print("Debug message")
	expected := "[DEBUG] Debug message\n"
	if out.String() != expected {
		t.Errorf("Expect is %q, but %q", expected, out.String())
	}

	out.Reset()

	dl.Printf("Debug %v %v", "message", "with format")
	expected = "[DEBUG] Debug message with format\n"
	if out.String() != expected {
		t.Errorf("Expect is %q, but %q", expected, out.String())
	}
}

func TestDebugEnvNotSpecified(t *testing.T) {
	os.Setenv("DEBUG", "")

	out := new(bytes.Buffer)
	dl := New(out)

	dl.Print("Debug message")
	expected := ""
	if out.String() != expected {
		t.Errorf("Expect is %q, but %q", expected, out.String())
	}

	out.Reset()

	dl.Printf("Debug %v %v", "message", "with format")
	expected = ""
	if out.String() != expected {
		t.Errorf("Expect is %q, but %q", expected, out.String())
	}
}

func TestWithFlag(t *testing.T) {
	os.Setenv("DEBUG", "1")

	out := new(bytes.Buffer)
	dl := New(out, Flag(log.Ldate))

	dl.Print("Debug message")

	year, month, day := time.Now().Date()
	expected := fmt.Sprintf("[DEBUG] %d/%02d/%02d Debug message\n", year, month, day)
	if out.String() != expected {
		t.Errorf("Expect is %q, but %q", expected, out.String())
	}
}

func TestMain(m *testing.M) {
	beforeEnv := os.Getenv("DEBUG")
	result := m.Run()
	os.Setenv("DEBUG", beforeEnv)
	os.Exit(result)
}
