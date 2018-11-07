# debuglog

debuglog is a logger for debugging. It outputs log only when `DEBUG` env is specified.

## Usage

```go
package main

import (
	"os"

	"github.com/y-yagi/debuglog"
)

func main() {
	logger := debuglog.New(os.Stdout)
	debuglog.Printf("Debug: %v\n", 1) // Nothing output

	os.Setenv("DEBUG", "1")

	logger = debuglog.New(os.Stdout)
	logger.Printf("Debug: %v\n", 1) // Log output
}
```
