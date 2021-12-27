package main

import (
	"fmt"
	"github.com/strategicpause/adventofcode2021/common"
	"strings"
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
	numSegments := 0
	common.SplitItr(input, '\n', func(s1 string) {
		i := strings.IndexByte(s1, '|')
		common.SplitItr(s1[i:], ' ', func(s string) {
			switch len(s) {
			case 2,3,4,7:
				numSegments += 1
			}
		})
	})
	return numSegments
}

func PartB(input string) int {
	total := 0
	common.SplitItr(input, '\n', func(s1 string) {
		i := strings.IndexByte(s1, '|')
		numMap := Learn(s1[0:i - 1])
		// At this point we have the map filled out and will be able 
		// to solve the remaining problem.
		total += Solve(s1[i + 2:], numMap)
	})
	return total
}

func Learn(s1 string) map[int]int {
	numMap := map[int]int {}
	var remaining []string
	var encoded4 string
	var encoded7 string
	// During the first pass, let's find all of the numbers that we know about.
	common.SplitItr(s1, ' ', func(s string) {
		switch len(s) {
		case 2:
			numMap[Hash(s)] = 1
		case 3:
			numMap[Hash(s)] = 7
			encoded7 = s
		case 4:
			numMap[Hash(s)] = 4
			encoded4 = s
		case 7:
			numMap[Hash(s)] = 8
		default:
			remaining = append(remaining, s)
		}			
	})
	// Work on remaining items in the second pass. We should be able to use the 4
	// and 7 to determine the remaining numbers.
	for _, r := range remaining {
		numCommon4 := common.NumCommon(encoded4, r)
		numCommon7 := common.NumCommon(encoded7, r)
		switch len(r) {
		case 5: // Detect 2, 3, or 5
			if numCommon4 == 3 && numCommon7 == 3 {
				numMap[Hash(r)] = 3
			} else if numCommon4 == 2 && numCommon7 == 2 {
				numMap[Hash(r)] = 2
			} else if numCommon4 == 3 && numCommon7 == 2 {
				numMap[Hash(r)] = 5
			}
		case 6: // Detect 0, 6, or 9
			if numCommon4 == 3 && numCommon7 == 3 {
				numMap[Hash(r)] = 0
			} else if numCommon4 == 3 && numCommon7 == 2 {
				numMap[Hash(r)] = 6
			} else if numCommon4 == 4 && numCommon7 == 3 {
				numMap[Hash(r)] = 9
			}
		}
	}
	return numMap
}

// Hash will return a product of all of the individual characters. Since multiplication is
// associative, this will return the same value for permutations of the same string.
func Hash(s string) int {
	product := 1
	for _, c := range s {
		product *= int(c)
	}
	return product
}

func Solve(s1 string, m map[int]int) int {
	answer := 0
	common.SplitItr(s1, ' ', func(s string) {
		answer *= 10
		answer += m[Hash(s)] 
	})
	return answer
}