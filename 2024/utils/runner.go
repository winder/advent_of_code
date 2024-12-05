package utils

import (
	"fmt"
)

func init() {
	days = make(map[int]Day)
}

var days map[int]Day

type RunHandler func() error

type Day struct {
	parts []RunHandler
}

func RegisterDay(day int, parts ...RunHandler) {
	if _, ok := days[day]; ok {
		panic(fmt.Sprintf("Duplicate registration for day %d", day))
	}
	days[day] = Day{
		parts: parts,
	}
}

func runOne(day Day) {
	all := onlyPart == 0

	for i, part := range day.parts {
		dayNum := i + 1
		fmt.Printf("Part %d\n", dayNum)
		if !all && onlyPart != dayNum {
			fmt.Println("** skipped **")
			continue
		}
		if err := part(); err != nil {
			fmt.Printf("** error: %v **\n", err)
		}
	}
}

func Run() {
	for i := 1; i <= len(days); i++ {
		if onlyDay != 0 && i != onlyDay {
			continue
		}
		if _, ok := days[i]; !ok {
			panic(fmt.Sprintf("Day %d is missing", i))
		}

		fmt.Println("==================")
		fmt.Printf("    %d\n", i)
		fmt.Println("==================")
		runOne(days[i])
		fmt.Println()
	}
}
