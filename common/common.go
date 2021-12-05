package common

import (
	"io/ioutil"
)

func ReadInput() (string, error) {
	content, err := ioutil.ReadFile("input.txt")
	return string(content), err
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