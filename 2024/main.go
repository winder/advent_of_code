package main

import (
	"fmt"
	"os"

	"github.com/winder/advent_of_code/2024/d01"
	"github.com/winder/advent_of_code/2024/d02"
	"github.com/winder/advent_of_code/2024/utils"
)

func main() {
	if err := utils.ValidateFlags(); err != nil {
		fmt.Printf("Invalid flags: %s\n", err)
		os.Exit(1)
	}

	d01.Run()
	d02.Run()
}
