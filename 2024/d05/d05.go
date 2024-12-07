package d05

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/winder/advent_of_code/2024/utils"
)

func init() {
	utils.RegisterDay(5, part, part)
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

			before[pg1] = append(before[pg1], pg2)
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
	/*
		remaining := make(map[int]struct{})
		for _, page := range book {
			remaining[page] = struct{}{}
		}

		for _, page := range book {
			// check for required pages
			for _, reqPage := range after[page] {
				if _, ok := remaining[reqPage]; ok {
					return false
				}
			}
			delete(remaining, page)
		}
	*/

	seen := make(map[int]struct{})
	for _, page := range book {
		for _, beforePage := range before[page] {
			if _, ok := seen[beforePage]; ok {
				// violation detected.
				return false
			}
		}
		seen[page] = struct{}{}
	}
	return true
}

func isOrderedOrReOrder(before map[int][]int, after map[int][]int, book []int) (bool, []int) {
	remaining := make(map[int]int)
	for i, page := range book {
		remaining[page] = i
	}

	reordered := false
	for i, _ := range book {
		swapped := true
		nextIdx := i

		// keep selecting from remaining pages until there are no violations.
		for swapped {
			swapped = false
			for _, cannotBeBeforePage := range after[book[nextIdx]] {
				if j, ok := remaining[cannotBeBeforePage]; ok {
					nextIdx = j
					swapped = true
				}
			}
		}

		// Perform the swap if needed.
		if i != nextIdx {
			reordered = true
			remaining[book[i]] = nextIdx
			book[i], book[nextIdx] = book[nextIdx], book[i]
		}

		delete(remaining, book[i])
	}

	return !reordered, book
}

func part(p int) error {
	filename := utils.InputFilename("d05/input")
	before, after, books, err := parse(filename)
	if err != nil {
		return err
	}

	sumOfOrdered := 0
	sumOfReordered := 0

	for _, book := range books {
		inOrder, book := isOrderedOrReOrder(before, after, book)

		mid := book[len(book)/2]
		if inOrder {
			sumOfOrdered += mid
		} else {
			sumOfReordered += mid
		}
	}

	switch p {
	case 1:
		fmt.Println(sumOfOrdered)
	case 2:
		fmt.Println(sumOfReordered)
	}

	return nil
}
