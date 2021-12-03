package main

import (
	"fmt"
	"io/ioutil"
	"strings"
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
			n = Atoi(input)
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
			n = Atoi(input)
			input = ""
		} else {
			n = Atoi(input[0:i])
			input = input[i+1:]
		}
		window = append(window, n)
		if len(window) > 3 {
			if Sum(window[0:3]) < Sum(window[1:4]) {
				increase++
			}
			window = window[1:4]
		}
	}
	return increase
}

func Atoi(str string) int {
	n := 0
	l := len(str)
	for i := 0; i < l; i++ {
		n = n * 10 + int(str[i] - '0')
	}
	return n
}

func Sum(window []int) int {
	sum := 0
	for _, n := range window {
		sum += n
	}
	return sum
}
