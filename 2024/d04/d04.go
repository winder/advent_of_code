package d04

import (
	"fmt"

	"github.com/winder/advent_of_code/2024/utils"
)

func init() {
	utils.RegisterDay(4, search(dfsP1), search(dfsP2))
}

func readGrid(filename string) ([][]byte, error) {
	var grid [][]byte
	rowlen := 0

	input, err := utils.Lines(filename)
	if err != nil {
		return nil, err
	}

	for line := range input {
		if rowlen == 0 {
			rowlen = len(line)
		}
		if len(line) != rowlen {
			return nil, fmt.Errorf("rows are different lengths %d != %d", rowlen, len(line))
		}
		grid = append(grid, []byte(line))
	}

	return grid, nil
}

func isXmas(x, m, a, s byte) bool {
	return x == 'X' && m == 'M' && a == 'A' && s == 'S'
}

// dfsP1 returns the number of 'xmas' matches that start at (x,y)
func dfsP1(grid [][]byte, x, y int) int {
	width := len(grid[0])
	height := len(grid)
	up := y >= 3
	down := y <= height-4
	right := x <= width-4
	left := x >= 3

	num := 0

	// up
	if up && isXmas(grid[y][x], grid[y-1][x], grid[y-2][x], grid[y-3][x]) {
		num++
	}
	// down
	if down && isXmas(grid[y][x], grid[y+1][x], grid[y+2][x], grid[y+3][x]) {
		num++
	}
	// right
	if right && isXmas(grid[y][x], grid[y][x+1], grid[y][x+2], grid[y][x+3]) {
		num++
	}
	// left
	if left && isXmas(grid[y][x], grid[y][x-1], grid[y][x-2], grid[y][x-3]) {
		num++
	}
	// up-right
	if up && right && isXmas(grid[y][x], grid[y-1][x+1], grid[y-2][x+2], grid[y-3][x+3]) {
		num++
	}
	// up-left
	if up && left && isXmas(grid[y][x], grid[y-1][x-1], grid[y-2][x-2], grid[y-3][x-3]) {
		num++
	}
	// up-right
	if down && right && isXmas(grid[y][x], grid[y+1][x+1], grid[y+2][x+2], grid[y+3][x+3]) {
		num++
	}
	// down-left
	if down && left && isXmas(grid[y][x], grid[y+1][x-1], grid[y+2][x-2], grid[y+3][x-3]) {
		num++
	}

	return num
}

func isMS(m, s byte) bool {
	return m == 'M' && s == 'S' || m == 'S' && s == 'M'
}

// dfsP2 returns the number of diagonal 'mas' matches that have A (x,y)
func dfsP2(grid [][]byte, x, y int) int {
	width := len(grid[0])
	height := len(grid)
	up := y >= 1
	down := y <= height-2
	right := x <= width-2
	left := x >= 1

	if !up || !down || !right || !left {
		return 0
	}

	num := 0
	if grid[x][y] == 'A' &&
		isMS(grid[x-1][y+1], grid[x+1][y-1]) &&
		isMS(grid[x-1][y-1], grid[x+1][y+1]) {
		num++
	}

	// greater than 1842
	return num
}

func search(dfs func(grid [][]byte, x, y int) int) func() error {
	return func() error {
		grid, err := readGrid(utils.InputFilename("d04/input"))
		if err != nil {
			return err
		}

		sum := 0
		for x := 0; x < len(grid[0]); x++ {
			for y := 0; y < len(grid); y++ {
				sum += dfs(grid, x, y)
			}
		}

		fmt.Println(sum)
		return nil
	}
}
