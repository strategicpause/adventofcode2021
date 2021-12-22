package main

import (
	"fmt"
	"github.com/strategicpause/adventofcode2021/common"
)

const (
	PartADays = 80
	PartBDays = 256
	MaxDays = 9
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
	fishes := common.SplitAtoi(input, ',')
	for day := 0; day < PartADays; day++ {
		n := len(fishes)
		for i := 0; i < n; i++ {
			if fishes[i] == 0 {
				fishes[i] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[i] -= 1
			}
		}
	}
	return len(fishes)
}

func PartB(input string) int {
	var fishes [MaxDays]int
	var newFishes [MaxDays]int
	common.SplitItr(input, ',', func(s string) {
		fishes[common.Atoi(s)] += 1
	})

	for day := 0; day < PartBDays; day++ {
		// Clear oldFishes slice
		for i := 0; i < MaxDays; i++ {
			newFishes[i] = 0
		}
		for i := MaxDays - 1; i >= 0 ; i-- {
			if i == 0 {
				newFishes[8] += fishes[0]
				newFishes[6] += fishes[0]
			} else {
				newFishes[i - 1] = fishes[i]
			}
		}
		fishes = newFishes
	}
	return common.Sum(fishes[:])
}