package wdx

import (
	"fmt"
)

// Log progress / details on the go
func Log(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
