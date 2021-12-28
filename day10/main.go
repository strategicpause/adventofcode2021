package main

import (
	"fmt"
	"github.com/strategicpause/adventofcode2021/common"
)

type Stack []rune

func (s *Stack) Push(c rune) {
	*s = append(*s, c)
}

func (s *Stack) Pop() rune {
	i := len(*s) - 1
	e := (*s)[i]
	*s = (*s)[:i]
	return e
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

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
	score := 0
	scores := map[rune]int {
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	closingChars := map[rune]rune {
		'}': '{',
		')': '(',
		']': '[',
		'>': '<',
	}

	common.SplitLines(input, func(line string) {
		stack := Stack {}
		for _, c := range line {
			switch c {
			case '{','(','[','<':
				stack.Push(c)
			case '}',')',']','>':
				if stack.Pop() != closingChars[c] {
					score += scores[c]
					break
				}
			}
		}
	})
	return score
}

func PartB(input string) int {
	charScores := map[rune]int {
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	closingChars := map[rune]rune {
		'}': '{',
		')': '(',
		']': '[',
		'>': '<',
	}

	var scores []int
	common.SplitLines(input, func(line string) {
		score := 0
		stack := Stack {}
		isCorrupted := false
		for _, c := range line {
			switch c {
			case '{','(','[','<':
				stack.Push(c)
			case '}',')',']','>':
				if stack.Pop() != closingChars[c] {
					isCorrupted = true
					break
				}
			}
		}
		if !isCorrupted {
			for !stack.IsEmpty() {
				c := stack.Pop()
				score = score * 5 + charScores[c]
			}
			scores = append(scores, score)
		}
	})
	return common.MiddleElement(scores)
}