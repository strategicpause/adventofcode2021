package common

import (
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func ReadInput() (string, error) {
	content, err := ioutil.ReadFile("input.txt")
	return string(content), err
}

func Atoi(str string) int {
	n := 0
	l := len(str)
	for i := 0; i < l; i++ {
		n = n * 10 + CharAtoi(str[i])
	}
	return n
}

func CharAtoi(c byte) int {
	return int(c - '0')
}

// SplitAtoi will split the given string by the given split character.
// Each of the resulting elements will then be converted to an integer.
func SplitAtoi(str string, splitChar byte) []int {
	var nums []int
	SplitItr(str, splitChar, func(s string) {
		nums = append(nums, Atoi(s))
	})
	return nums
}

func SplitLines(str string, f func(string)) {
	SplitItr(str, '\n', f)
}

// SplitItr will split the given string by the given split character.
// The resulting elements will passed as a parameter to the given
// function.
func SplitItr(str string, splitChar byte, f func(string)) {
	for str != "" {
		i := strings.IndexByte(str, splitChar)
		if i == -1 {
			f(str)
			str = ""
		} else {
			f(str[0:i])
			str = str[i+1:]
		}
	}
}

func Sum(window []int) int {
	sum := 0
	for _, n := range window {
		sum += n
	}
	return sum
}

func Abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func Max(n, m int) int {
	if n > m {
		return n
	}
	return m
}

func GetLowHigh(nums []int) (int, int) {
	high := 0
	low := math.MaxInt
	for _, num := range nums {
		if num > high {
			high = num
		} else if num < low {
			low = num
		}
	}
	return low, high
}

// Sumation returns the sum of numbers from 1 to n (inclusive).
func Summation(n int) int {
	return int(float64(n) / 2.0 * float64(1 + n))
}

// NumCommon returns the number of characters that are in common between s1 and s2.
func NumCommon(s1, s2 string) int {
	runes := map[rune]int {}
	for _, c := range s1 {
		runes[c] = 1
	}
	common := 0
	for _, c := range s2 {
		if _, ok := runes[c]; ok {
			common += 1
		}
	}
	return common
}

func MiddleElement(n []int) int {
	sort.Ints(n)
	return n[len(n)/2]
}
