package main

import (
	"fmt"
	"os"

	_ "github.com/winder/advent_of_code/2024/d01"
	_ "github.com/winder/advent_of_code/2024/d02"
	_ "github.com/winder/advent_of_code/2024/d03"
	_ "github.com/winder/advent_of_code/2024/d04"
	"github.com/winder/advent_of_code/2024/utils"
)

func main() {
	if err := utils.ValidateFlags(); err != nil {
		fmt.Printf("Invalid flags: %s\n", err)
		os.Exit(1)
	}

	utils.Run()
}
