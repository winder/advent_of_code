package d02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/winder/advent_of_code/2024/utils"
)

func init() {
	test := func(skip int) func() error {
		return func() error {
			var wg sync.WaitGroup
			wg.Add(1)
			chan1 := make(chan string)
			go func() {
				defer wg.Done()
				solution(chan1, skip)
			}()
			scanInput(utils.InputFilename("d02/input"), chan1)
			wg.Wait()
			return nil
		}
	}
	utils.RegisterDay(2, test(0), test(1))
}

func scanInput(filepath string, results ...chan<- string) error {
	for _, part := range results {
		defer close(part)
	}

	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("unable to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, part := range results {
			part <- scanner.Text()
		}
	}
	return scanner.Err()
}

func diff(a, b int) int {
	if b > a {
		a, b = b, a
	}
	return a - b
}

func isAsc(nums []int) bool {
	dsc := 0
	asc := 0
	first := nums[0]
	for _, num := range nums[1:] {
		if first < num {
			asc++
		} else if num < first {
			dsc++
		}
	}
	return asc > dsc
}

func isValid(nums []int, skips int) bool {
	ascending := isAsc(nums)
	badLevel := 0
	last := nums[0]

	for _, num := range nums[1:] {

		if badLevel > skips {
			utils.Debugf("%v: Unsafe (bad: %d, skips: %d)\n", nums, badLevel, skips)
			return false
		}

		if last == num {
			badLevel++
			continue
		} else if d := diff(last, num); d < 1 || d > 3 {
			badLevel++
			continue
		} else if ascending && last > num {
			badLevel++
			continue
		} else if !ascending && last < num {
			badLevel++
			continue
		}

		// only update if num is good.
		last = num
	}

	if badLevel > skips {
		utils.Debugf("%v: Unsafe (bad: %d, skips: %d)\n", nums, badLevel, skips)
		return false
	}

	utils.Debugf("%v: Safe (bad: %d, skips: %d)\n", nums, badLevel, skips)
	return true
}

func solution(lines <-chan string, skips int) {
	validCount := 0
	for line := range lines {
		numstrs := strings.Fields(line)
		var nums []int
		for _, numstr := range numstrs {
			num, err := strconv.Atoi(numstr)
			if err != nil {
				fmt.Println("parse error: ", err)
				os.Exit(1)
			}
			nums = append(nums, num)
		}

		if isValid(nums, skips) || isValid(nums[1:], skips-1) {
			validCount++
		}
	}

	fmt.Printf("Valid reports (skips %d): %d\n", skips, validCount)
}
