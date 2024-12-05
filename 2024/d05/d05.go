package d05

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/winder/advent_of_code/2024/utils"
)

func init() {
	utils.RegisterDay(5, part1, part2)
}

func parse(filepath string) (map[int][]int, map[int][]int, [][]int, error) {
	input, err := utils.Lines(filepath)
	if err != nil {
		return nil, nil, nil, err
	}

	before := make(map[int][]int)
	after := make(map[int][]int)
	pt2 := false
	books := make([][]int, 0)
	for ruleOrBook := range input {
		if !pt2 {
			if ruleOrBook == "" {
				// done with rule section
				//break
				pt2 = true
				continue
			}
			var pg1 int
			var pg2 int
			_, err := fmt.Sscanf(ruleOrBook, "%d|%d", &pg1, &pg2)
			if err != nil {
				return nil, nil, nil, fmt.Errorf("bad rule: %s", ruleOrBook)
			}

			before[pg1] = append(after[pg1], pg2)
			after[pg2] = append(after[pg2], pg1)
		} else {
			var book []int
			pages := strings.Split(ruleOrBook, ",")
			for _, page := range pages {
				p, err := strconv.Atoi(page)
				if err != nil {
					return nil, nil, nil, fmt.Errorf("bad page %s in rawBook %s", page, ruleOrBook)
				}
				book = append(book, p)
			}
			books = append(books, book)
		}
	}

	return before, after, books, nil
}

func isOrdered(before map[int][]int, after map[int][]int, book []int) bool {
	remaining := make(map[int]struct{})
	for _, page := range book {
		remaining[page] = struct{}{}
	}

	pageSet := make(map[int]struct{})
	for _, page := range book {
		// check for required pages
		for _, reqPage := range after[page] {
			if _, ok := remaining[reqPage]; ok {
				return false
			}
		}
		delete(remaining, page)

		pageSet[page] = struct{}{}
	}
	return true
}

func part1() error {
	filename := utils.InputFilename("d05/input")
	before, after, books, err := parse(filename)
	if err != nil {
		return err
	}

	sumOfMiddles := 0

	for _, book := range books {
		if isOrdered(before, after, book) {
			sumOfMiddles += book[len(book)/2]
		}
	}

	fmt.Println(sumOfMiddles)

	return nil
}

func reOrder(before map[int][]int, after map[int][]int, book []int) []int {
	remaining := make(map[int]struct{})
	for _, page := range book {
		remaining[page] = struct{}{}
	}

	return nil
}

func part2() error {
	filename := utils.InputFilename("d05/input")
	before, after, books, err := parse(filename)
	if err != nil {
		return err
	}

	sumOfMiddles := 0

	for _, book := range books {
		if isOrdered(before, after, book) {
			continue
			//sumOfMiddles += book[len(book)/2]
		}
	}

	fmt.Println(sumOfMiddles)

	return nil
}
