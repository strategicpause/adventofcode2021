package main

import (
	"fmt"
	"strings"
	"github.com/strategicpause/adventofcode2021/common"
)

const (
	BoardSize = 5
)

type Board struct {
	cols [BoardSize]int
	rows [BoardSize]int
	spaces map[string]*Space
}

func NewBoard(input string) *Board {
	board := Board{
		spaces: make(map[string]*Space),
	}
	lines := strings.Split(input, "\n")
	row := 0
	for col, line := range lines {
		values := strings.Split(line, " ")
		for _, value := range values {
			value = strings.TrimSpace(value)
			if value != "" {
				board.spaces[value] = NewSpace(value, row, col)
				row++	
			}
		}
		row = 0
	}
	return &board
}

type Space struct {
	value string
	row int
	col int
	marked bool
}

func NewSpace(value string, row, col int) *Space {
	return &Space {
		value: value,
		row: row,
		col: col,
		marked: false,
	}
}

func (b *Board) MarkSpace(value string) bool {
	if space, ok := b.spaces[value]; ok {
		space.marked = true
		b.rows[space.row]++
		if b.rows[space.row] == BoardSize {
			return true
		}
		b.cols[space.col]++
		if b.cols[space.col] == BoardSize {
			return true
		}
	}
	return false
}

func (b *Board) GetScore(value string) int {
	sum := 0
	for _, space := range b.spaces {
		if !space.marked {
			sum += common.Atoi(space.value)
		}
	}
	return sum * common.Atoi(value)
}

type Game struct {
	moves []string
	boards []*Board
}

func (g *Game) Play() int {
	for _, move := range g.moves {
		for _, board := range g.boards {
			if board.MarkSpace(move) {
				return board.GetScore(move)
			}
		}
	}
	return -1
}

func NewGame(input string) *Game {
	game := Game{}

	i := strings.Index(input, "\n")
	game.moves = strings.Split(input[0:i],",")
	input = input[i:]

	for len(input) > 0 {
		i = strings.Index(input, "\n\n")
		if i == -1 {
			game.boards = append(game.boards, NewBoard(input))
			input = ""
		} else {
			game.boards = append(game.boards, NewBoard(input[0:i]))
			input = strings.TrimSpace(input[i:])	
		}
		
	}
	return &game
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
	game := NewGame(input)
	return game.Play()
}

func PartB(input string) int {
	return 0
}