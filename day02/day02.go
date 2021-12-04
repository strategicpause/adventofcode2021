package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/strategicpause/adventofcode2021/common"
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
	lines := strings.Split(input, "\n")
	depth := 0
	horizontalPos := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		i := strings.IndexByte(line, ' ')
		command := line[0:i]
		value := common.Atoi(line[i + 1:])
		switch command {
		case "forward":
			horizontalPos += value
		case "down":
			depth += value
		case "up":
			depth -= value
		default:
			fmt.Println(command)
		}
	}
	return depth * horizontalPos
}

func PartB(input string) int {
	lines := strings.Split(input, "\n")
	aim := 0
	depth := 0
	horizontalPos := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		i := strings.IndexByte(line, ' ')
		command := line[0:i]
		value := common.Atoi(line[i + 1:])
		switch command {
		case "forward":
			horizontalPos += value
			depth -= aim * value
		case "down":
			aim -= value
		case "up":
			aim += value
		default:
			fmt.Println(command)
		}
	}
	return depth * horizontalPos
}