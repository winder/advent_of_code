package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/winder/advent_of_code/2024/d01"
	"github.com/winder/advent_of_code/2024/d02"
	"github.com/winder/advent_of_code/2024/utils"
)

func main() {
	flag.BoolVar(&utils.Verbose, "verbose", false, "verbose output")
	flag.BoolVar(&utils.OnlyPart1, "only-part-1", false, "only run part 1")
	flag.BoolVar(&utils.OnlyPart2, "only-part-2", false, "only run part 2")
	flag.Parse()

	if utils.OnlyPart1 && utils.OnlyPart2 {
		fmt.Println("Do not provide both -only-part-1 and -only-part-2")
		os.Exit(1)
	}

	d01.Run()
	d02.Run()
}
