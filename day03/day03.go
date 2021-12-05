package main

import (
	"fmt"
	"github.com/strategicpause/adventofcode2021/common"
	"strconv"
	"strings"
)

const (

	MaxLen = 12
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
	// This array indicates the number of 1 bits in each column.
	var bits [MaxLen]int
	i := 0
	numLines := 0
	numBits := 0
	for _, c := range input {
		if c == '1' { 
			bits[i]++
			i++
		} else if c == '0' {
			// Just move the pointer
			i++
		} else if c == '\n' {
			numBits = i
			i = 0
			numLines++
		}
	}
	// Account for the last line which does not have a new line character
	numLines++
	gamma := 0
	for i := 0; i < numBits; i++ {
		gamma = gamma << 1
		if bits[i] >= (numLines / 2)  {
			gamma |= 1
		}
	}
	// Epsilon can be found by NOTing the gamma value. However, this will
	// result in all of the other 0 bits getting set to 1, so you need to 
	// create a mask which will only give us the bits that make up the 
	// actual epsilon value.
	epsilon := ^gamma & ((1 << numBits) - 1)
	return gamma * epsilon
}

func PartB(input string) int {
	ratings := strings.Split(input, "\n")
	numBits := len(ratings[0])
	oRatings := ratings
	co2Ratings := ratings
	for i := 0; i < numBits; i++ {
		oRatings = FilterRatings(oRatings, i, OxygenFilterFunc)
		co2Ratings = FilterRatings(co2Ratings, i, CarbonDioxideFilterFunc)
	}
	return BinStr2Int(oRatings[0]) * BinStr2Int(co2Ratings[0])
}

func OxygenFilterFunc(numBits, numLines int) byte {
	mostCommon := byte('0')
	if float32(numBits) >= float32(numLines) / 2.0 {
		mostCommon = byte('1')
	}
	return mostCommon
}

func CarbonDioxideFilterFunc(numBits, numLines int) byte {
	leastCommon := byte('1')
	if float32(numBits) >= float32(numLines) / 2.0 {
		leastCommon = byte('0')
	}
	return leastCommon
}

func CalcBitsAtIndex(ratings []string, index int) int {
	numBits := 0
	for _, rating := range ratings {
		if rating[index] == '1' {
			numBits++
		}
	}
	return numBits
}

func FilterRatings(ratings []string, index int, filterFunc func(int, int) byte) []string {
	if len(ratings) == 1 {
		return ratings
	}
	numBits := CalcBitsAtIndex(ratings, index)
	numLines := len(ratings)
	var newRatings [MaxLen]string
	i := 0
	filterValue := filterFunc(numBits, numLines)
	for _, rating := range ratings {
		if rating[index] == filterValue {
			newRatings[i] = rating
			i++
		}
	}
	return newRatings[0:i]
}

func BinStr2Int(binStr string) int {
	i, _ := strconv.ParseInt(binStr, 2, 64)
	return int(i)
}