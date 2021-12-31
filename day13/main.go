package main

import (
	"fmt"
	"strings"
	"github.com/strategicpause/adventofcode2021/common"
)

const (
	XAxis = 'x'
	YAxis = 'y'
	Width = 40
	Height = 6
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

type Dot struct {
	X int
	Y int
}

func (d *Dot) Hash() int {
	return d.X * 5000 + d.Y
}

type Fold struct {
	Axis byte
	Value int
}

type Instructions struct {
	Dots []*Dot
	Folds []*Fold
}

func PartA(input string) int {
	instructions := ReadInstructions(input)
	PerformFold(instructions.Folds[0], instructions.Dots)
	return CountDots(instructions)
}

func ReadInstructions(input string) *Instructions {
	dotMode := true
	instructions := Instructions {
		Dots: []*Dot{},
		Folds: []*Fold{},
	}

	common.SplitLines(input, func(line string) {
		if line == "" {
			dotMode = false
			return
		}
		if dotMode {
			i := strings.IndexByte(line, ',')
			x := common.Atoi(line[:i])
			y := common.Atoi(line[i+1:])
			instructions.Dots = append(instructions.Dots, &Dot{X:x,Y:y})
		} else {
			fold := Fold {
				Axis: line[11],
				Value: common.Atoi(line[13:]),
			}
			instructions.Folds = append(instructions.Folds, &fold)
		}
	})

	return &instructions
}

func PerformFolds(inst *Instructions) {
	for _, fold := range inst.Folds {
		PerformFold(fold, inst.Dots)
	}
}

func PerformFold(fold *Fold, dots []*Dot) {
	for _, dot := range dots {
		if fold.Axis == XAxis && dot.X > fold.Value {
			dot.X = 2 * fold.Value - dot.X
		} else if fold.Axis == YAxis && dot.Y > fold.Value {
			dot.Y = 2 * fold.Value - dot.Y
		}
	}
}

func CountDots(inst *Instructions) int {
	distinctDots := map[int]bool{}
	for _, dot := range inst.Dots {
		distinctDots[dot.Hash()] = true
	}
	return len(distinctDots)
}

func PartB(input string) int {
	instructions := ReadInstructions(input)
	PerformFolds(instructions)
	Print(instructions.Dots)
	return 1
}

func Print(dots []*Dot) {
	output := [Height][Width]rune {}
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			output[i][j] = '.'
		}
	}
	for _, dot := range dots {
		output[dot.Y][dot.X] = '*'
	}
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			fmt.Print(string(output[i][j]))
		}
		fmt.Println()
	}
}