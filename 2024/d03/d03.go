package d03

import (
	"fmt"
	"iter"
	"regexp"
	"strconv"

	"github.com/winder/advent_of_code/2024/utils"
)

func init() {
	utils.RegisterDay(3, part1, part2)
}

type op struct {
	op     string
	v1, v2 int
}

func (m op) multiply() int {
	return m.v1 * m.v2
}

func makeOp(parts []string) (op, error) {
	if len(parts) != 5 {
		return op{}, fmt.Errorf("match has 5 parts")
	}

	var err error
	// janky, but regexp should make sure there is only one
	o := op{op: fmt.Sprintf("%s%s%s", parts[0], parts[3], parts[4])}

	if o.op != "mul" {
		return o, nil
	}

	if len(parts) < 3 {
		return op{}, fmt.Errorf("mul needs at least 3 parts")
	}

	o.v1, err = strconv.Atoi(parts[1])
	if err != nil {
		return op{}, fmt.Errorf("non-integer part1: %s, %s", parts[1], parts)
	}
	o.v2, err = strconv.Atoi(parts[2])
	if err != nil {
		return op{}, fmt.Errorf("non-integer part2: %s, %s", parts[2], parts)
	}
	return o, nil
}

func tokenize(lines iter.Seq[string]) iter.Seq[op] {
	var re = regexp.MustCompile(`(mul)\((\d+),(\d+)\)|(do)\(\)|(don't)\(\)`)
	return func(yield func(m op) bool) {
		for line := range lines {
			for _, match := range re.FindAllStringSubmatch(line, -1) {
				o, err := makeOp(match[1:])
				if err != nil {
					panic(fmt.Sprintf("Should not happen unless regex is wrong: %s", err))
				}
				if !yield(o) {
					return
				}
			}
		}
	}
}

func part1() error {
	filename := utils.InputFilename("d03/input")
	input, err := utils.Lines(filename)
	if err != nil {
		return err
	}

	sum := 0
	for m := range tokenize(input) {
		if m.op == "mul" {
			sum += m.multiply()
		}
	}

	fmt.Println("sum: ", sum)
	return nil
}

func part2() error {
	filename := utils.InputFilename("d03/input")
	input, err := utils.Lines(filename)
	if err != nil {
		return err
	}

	sum := 0
	enabled := true
	for m := range tokenize(input) {
		switch m.op {
		case "mul":
			if enabled {
				sum += m.multiply()
			}
		case "do":
			enabled = true
		case "don't":
			enabled = false
		}
	}

	fmt.Println("sum: ", sum)
	return nil
}
