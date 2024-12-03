package utils

import (
	"flag"
	"fmt"
)

func init() {
	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.IntVar(&onlyDay, "only-day", 0, "only run the given day")
	flag.BoolVar(&onlyPart1, "only-part-1", false, "only run part 1")
	flag.BoolVar(&onlyPart2, "only-part-2", false, "only run part 2")
	flag.StringVar(&inputFile, "input-file", "", "override input filename.")
}

var verbose = false
var onlyDay = 0
var onlyPart1 = true
var onlyPart2 = true
var inputFile = ""

func ValidateFlags() error {
	flag.Parse()

	if onlyPart1 && onlyPart2 {
		return fmt.Errorf("do not provide both -only-part-1 and -only-part-2")
	}

	if onlyDay != 0 {
		if _, ok := days[onlyDay]; !ok {
			return fmt.Errorf("provided day %d is not registered", onlyDay)
		}
	}

	return nil
}

func InputFilename(defaultValue string) string {
	if inputFile == "" {
		return defaultValue
	}
	return inputFile
}

func Debugf(format string, args ...any) {
	if verbose {
		fmt.Printf(format, args...)
	}
}
