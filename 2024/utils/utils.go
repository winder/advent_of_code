package utils

import (
	"bufio"
	"flag"
	"fmt"
	"iter"
	"os"
)

func init() {
	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.IntVar(&onlyDay, "only-day", 0, "only run the given day")
	flag.IntVar(&onlyPart, "only-part", 0, "only run given part")
	flag.StringVar(&inputFile, "input-file", "", "override input filename.")
}

var verbose = false
var onlyDay = 0
var onlyPart = 0
var inputFile = ""

func ValidateFlags() error {
	flag.Parse()

	if onlyDay != 0 {
		if _, ok := days[onlyDay]; !ok {
			return fmt.Errorf("provided day %d is not registered", onlyDay)
		}
	}

	return nil
}

func InputFilename(defaultFile string) string {
	if inputFile == "" {
		return defaultFile
	}
	return inputFile
}

func Lines(defaultFile string) (iter.Seq[string], error) {
	file, err := os.Open(InputFilename(defaultFile))
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	scanner := bufio.NewScanner(file)

	seq := func(yield func(string) bool) {
		defer file.Close()
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				return
			}
		}
	}
	return seq, scanner.Err()
}

func Debugf(format string, args ...any) {
	if verbose {
		fmt.Printf(format, args...)
	}
}
