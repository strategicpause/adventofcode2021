package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/strategicpause/adventofcode2021/common"
)

const (
	WindowSize = 3
)

func main() {
	input, err := ReadInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	answer := PartA(input)
	fmt.Println(answer)
	answer = PartB(input)
	fmt.Println(answer)
}

func ReadInput() (string, error) {
	content, err := ioutil.ReadFile("input.txt")
	return string(content), err
}

func PartA(input string) int {
	prev := -1
	increase := 0
	for len(input) > 0 {
		i := strings.IndexByte(input, '\n')
		n := 0
		if i == -1 {
			n = common.Atoi(input)
			input = ""
		} else {
			n = Atoi(input[0:i])
			input = input[i+1:]
		}
		if prev != -1 && n > prev {
			increase++
		}
		prev = n
	}
	return increase
}

func PartB(input string) int {
	var window []int
	increase := 0
	for len(input) > 0 {
		i := strings.IndexByte(input, '\n')
		n := 0
		if i == -1 {
			n = common.Atoi(input)
			input = ""
		} else {
			n = common.Atoi(input[0:i])
			input = input[i+1:]
		}
		window = append(window, n)
		if len(window) > 3 {
			if common.Sum(window[0:3]) < common.Sum(window[1:4]) {
				increase++
			}
			window = window[1:4]
		}
	}
	return increase
}
