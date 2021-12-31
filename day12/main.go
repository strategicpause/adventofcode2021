package main

import (
	"fmt"
	"strings"
	"unicode"
	"github.com/strategicpause/adventofcode2021/common"
)

const (
	Start = "start"
	End = "end"
)

type Cave struct {
	Neighbors []string
	IsSmallCave bool
	Visited bool
}

type Graph map[string]*Cave

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
	graph := ReadGraph(input)
	return FindPaths(Start, graph)
}

func ReadGraph(input string) Graph {
	graph := Graph{}
	common.SplitLines(input, func(line string) {
		i := strings.IndexByte(line, '-')
		start := line[:i]
		end := line[i+1:]
		if _, ok := graph[start]; !ok {
			isSmallCave := unicode.IsLower(rune(start[0]))
			graph[start] = &Cave{Neighbors: []string{}, IsSmallCave: isSmallCave, Visited: false}
		}
		if _, ok := graph[end]; !ok {
			isSmallCave := unicode.IsLower(rune(end[0]))
			graph[end] = &Cave{Neighbors: []string{}, IsSmallCave: isSmallCave, Visited: false}
		}
		graph[start].Neighbors = append(graph[start].Neighbors, end)
		graph[end].Neighbors = append(graph[end].Neighbors, start)
	})
	return graph
}

func FindPaths(current string, graph Graph) int {
	if current == End {
		return 1
	}
	numPaths := 0
	cave := graph[current]
	cave.Visited = cave.IsSmallCave
	for _, neighbor := range cave.Neighbors {
		if !graph[neighbor].Visited {
			numPaths += FindPaths(neighbor, graph)
		}
	}
	cave.Visited = false
	return numPaths
}

func PartB(input string) int {
	graph := ReadGraph(input)
	return FindPaths2(Start, map[string]int{}, graph)
}

func FindPaths2(current string, visitedSmallCaves map[string]int, graph Graph) int {
	if current == End {
		return 1
	}
	
	cave := graph[current]
	if cave.IsSmallCave {
		hasTwoCaves := HasTwoCaves(visitedSmallCaves)
		// If we already have two caves and we're visiting a cave that
		// we have already visited, then return immediately.
		if hasTwoCaves && visitedSmallCaves[current] > 0 {
			return 0
		}
		visitedSmallCaves[current] += 1
		// Consider a cave visited if we're visiting it for the first time
		// and we have already visited the same cave twice. Otherwise 
		// consider it to be visited if this is the first cave we're 
		// visiting twice.
		cave.Visited = hasTwoCaves || visitedSmallCaves[current] == 2
	}

	numPaths := 0
	for _, neighbor := range cave.Neighbors {
		if neighbor != Start && !graph[neighbor].Visited {
			numPaths += FindPaths2(neighbor, visitedSmallCaves, graph)
		}
	}
	cave.Visited = false
	if cave.IsSmallCave {
		visitedSmallCaves[current] -= 1
	}
	return numPaths
}

func HasTwoCaves(visitedSmallCaves map[string]int) bool {
	for _, visits := range visitedSmallCaves {
		if visits == 2 {
			return true
		}
	}
	return false
}