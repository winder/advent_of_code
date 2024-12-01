package d01

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput(filepath string) ([]int, []int, error) {
	var l1 []int
	var l2 []int

	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines := strings.Split(line, "   ")
		if len(lines) != 2 {
			return nil, nil, fmt.Errorf("unexpected input line: %s", line)
		}

		i1, err := strconv.Atoi(lines[0])
		if err != nil {
			return nil, nil, fmt.Errorf("unable to parse i1: %w", err)
		}
		l1 = append(l1, i1)

		i2, err := strconv.Atoi(lines[1])
		if err != nil {
			return nil, nil, fmt.Errorf("unable to parse i2: %w", err)
		}
		l2 = append(l2, i2)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("scanner error: %w", err)
	}

	return l1, l2, nil
}

func part1(l1 []int, l2 []int) int {
	sum := 0
	for i, _ := range l1 {
		p1 := l1[i]
		p2 := l2[i]
		if p2 > p1 {
			p1, p2 = p2, p1
		}
		sum += (p1 - p2)
	}
	return sum
}

func part2(l1 []int, l2 []int) int {
	s2 := make(map[int]int)
	for _, v := range l2 {
		s2[v] = s2[v] + 1
	}
	sim := 0
	for _, v := range l1 {
		sim += v * s2[v]
	}
	return sim
}

func D1() {
	l1, l2, err := readInput("d01/1.input")
	if err != nil {
		fmt.Printf("Input error: %s\n", err)
		os.Exit(1)
	}
	sort.Ints(l1)
	sort.Ints(l2)

	if len(l1) != len(l2) {
		fmt.Printf("Unexpected inputs, expected l1 == l2, got: %d != %d\n", len(l1), len(l2))
		os.Exit(1)
	}

	fmt.Println("Part1:", part1(l1, l2))
	fmt.Println("Part2:", part2(l1, l2))
}
