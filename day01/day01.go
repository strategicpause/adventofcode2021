package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ReadInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	answer := PartA(input)
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

func Atoi(str string) int {
	multiplier := 1
	n := 0
	l := len(str)
	for i := 0; i < l; i++ {
		n += int(str[l - i - 1] - '0') * multiplier
		multiplier *= 10
	}
	return n
}
