package main

import (
	"fmt"
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
	return 0
}

func PartB(input string) int {
	return 0
}