package main

import (
	"fmt"	
	"strings"
	"github.com/strategicpause/adventofcode2021/common"
)

type Line struct {
	x1, y1 int
	x2, y2 int
}

func (l *Line) IsNonDiag() bool {
	return l.x1 == l.x2 || l.y1 == l.y2
}

func (l *Line) GetPoints() []*Point {
	iterations := common.Max(common.Abs(l.y2 - l.y1), common.Abs(l.x2 - l.x1)) + 1
	xItr := 0
	if l.x1 > l.x2 {
		xItr = -1 
	} else if l.x1 < l.x2 {
		xItr = 1
	}
	yItr := 0
	if l.y1 > l.y2 {
		yItr = -1 
	} else if l.y1 < l.y2 {
		yItr = 1
	}
	var points []*Point
	for i := 0; i < iterations; i++ {
		point := Point {
			x: l.x1 + xItr * i,
			y: l.y1 + yItr * i,
		}
		points = append(points, &point)
	}
	return points
}

type Point struct {
	x, y int
}

func (p *Point) GetValue() int {
	return p.x * 1000 + p.y
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
	inputLines := strings.Split(strings.TrimSpace(input), "\n")
	var lines []*Line
	for _, inputLine := range inputLines {
		line := getLine(inputLine)
		if line.IsNonDiag() {
			lines = append(lines, line)
		}
	}
	return GetNumIntersections(lines)
}

func getLine(line string) *Line {
	// Parse x1
	i := strings.IndexByte(line, ',')
	x1 := common.Atoi(line[0:i])
	line = line[i + 1:]
	// Parse y1
	i = strings.IndexByte(line, ' ')
	y1 := common.Atoi(line[0:i])
	line = line[i + 1:]
	// Parse x2
	i = strings.IndexByte(line, ',')
	x2 := common.Atoi(line[3:i])
	line = line[i + 1:]
	// Parse y2
	y2 := common.Atoi(line)

	return &Line {
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,
	}
}

func GetNumIntersections(lines []*Line) int {
	intersections := map[int]int {}
	for _, line := range lines {
		points := line.GetPoints()
		for _, point := range points {
			key := point.GetValue()
			intersections[key] += 1
		}
	}
	numIntersections := 0
	for _, value := range intersections {
		if value > 1 {
			numIntersections += 1
		}
	}

	return numIntersections
}

func PartB(input string) int {
	inputLines := strings.Split(strings.TrimSpace(input), "\n")
	var lines []*Line
	for _, inputLine := range inputLines {
		line := getLine(inputLine)
		lines = append(lines, line)
	}
	return GetNumIntersections(lines)
}