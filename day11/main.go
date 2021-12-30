package main

import (
	"fmt"
	"github.com/strategicpause/adventofcode2021/common"
)

const (
	Steps = 100
	GridSize = 10
	X = 0
	Y = 1
)

func main() {
	input, err := common.ReadInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	answer := PartA(input)
	fmt.Println(answer)
	answer = PartB(input)
	fmt.Println(answer)
}

var Offsets [][]int = [][]int {
    {-1, -1}, {0, -1}, {1, -1},
    {-1, 0}, {1, 0},
    {-1, 1}, {0, 1}, {1, 1},
}

type SquidGrid [GridSize][GridSize]int

func PartA(input string) int {
	grid := ReadGrid(input)
	flashes := 0
	for i := 0; i < Steps; i++ {
		flashes += RunStep(grid)
	}
	return flashes
}

func ReadGrid(input string) *SquidGrid {
	grid := SquidGrid{}
	y := 0
	common.SplitLines(input, func(s string) {
		for x := 0; x < GridSize; x++ {
			grid[y][x] = common.CharAtoi(s[x])
		}
		y += 1
	})
	return &grid
}

func RunStep(grid *SquidGrid) int {
	flashes := 0
	// Increase the energy of each squid by 1
	for y := 0; y < GridSize; y++ {
		for x := 0; x < GridSize; x++ {
			grid[y][x] += 1
		}
	}
	for y := 0; y < GridSize; y++ {
		for x := 0; x < GridSize; x++ {
			if grid[y][x] > 9 {
				flashes += Flash(x, y, grid)
			}
		}
	}
	
	return flashes
}

func Flash(x, y int, grid *SquidGrid) int {
	grid[y][x] = 0
	flashes := 1
	for _, offset := range Offsets {
		newX := x + offset[X]
		newY := y + offset[Y]
		if newX >= 0 && newY >= 0 && newX < GridSize && newY < GridSize {
			if grid[newY][newX] == 0 || grid[newY][newX] > 9 {
				continue
			}
			grid[newY][newX] += 1
			if grid[newY][newX] > 9 {
				flashes += Flash(newX, newY, grid)
			}
		}
	}
	return flashes
}

func PartB(input string) int {
	grid := ReadGrid(input)
	steps := 0
	for {
		RunStep(grid)
		steps += 1
		if CheckGrid(grid) {
			return steps
		}
	}
	return 0
}

func CheckGrid(grid *SquidGrid) bool {
	for y := 0; y < GridSize; y++ {
		for x := 0; x < GridSize; x++ {
			if grid[y][x] != 0 {
				return false
			}
		}
	}
	return true
}