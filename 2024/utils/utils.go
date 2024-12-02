package utils

import (
	"fmt"
)

var Verbose = false
var OnlyPart1 = false
var OnlyPart2 = false

func Debugf(format string, args ...any) {
	if Verbose {
		fmt.Printf(format, args...)
	}
}
