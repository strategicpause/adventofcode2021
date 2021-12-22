package main

import (
	"fmt"
	"math"
	"github.com/strategicpause/adventofcode2021/common"
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

func PartA(input string) int {
	positions := common.SplitAtoi(input, ',')
	low, high := common.GetLowHigh(positions)
	lowSum := math.MaxInt
	for i := low; i <= high; i++ {
		sum := 0
		for _, position := range positions {
			sum += common.Abs(i - position)
		}
		if sum < lowSum {
			lowSum = sum
		}
	}
	return lowSum
}

func PartB(input string) int {
	positions := common.SplitAtoi(input, ',')
	low, high := common.GetLowHigh(positions)
	lowSum := math.MaxInt
	for i := low; i <= high; i++ {
		sum := 0
		for _, position := range positions {
			sum += common.Summation(common.Abs(i - position))
		}
		if sum < lowSum {
			lowSum = sum
		}
	}
	return lowSum
}