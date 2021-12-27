package main

import (
	"container/heap"
	"fmt"
	"github.com/strategicpause/adventofcode2021/common"
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

type Entry struct {
	value int
	x int
	y int
	explored bool
}

type Queue struct {
	entries []*Entry
}

func (q *Queue) Push(e *Entry) {
	q.entries = append(q.entries, e)
}

func (q *Queue) Pop() *Entry {
	e := q.entries[0]
	q.entries = q.entries[1:]
	return e
}

func (q *Queue) IsEmpty() bool {
	return len(q.entries) == 0
}

func NewQueue() *Queue {
	return &Queue {
		entries: []*Entry{},
	}
}

// Shamelessly stolen from StackOverflow
type MaxIntHeap []int

func (h MaxIntHeap) Len() int { 
	return len(h) 
}

func (h MaxIntHeap) Less(i, j int) bool {
 return h[i] > h[j] 
}

func (h MaxIntHeap) Swap(i, j int) { 
	h[i], h[j] = h[j], h[i] 
}

func (h *MaxIntHeap) Push(x interface {}) {
	*h = append(*h, x.(int))
}

func (h *MaxIntHeap) Pop() interface {} {
	old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func NewMaxIntHeap() *MaxIntHeap {
	return &MaxIntHeap{}
}

func (e *Entry) GetNeighbors() []int {
	var neighbors [4]int
	// Add Top
	neighbors[0] = GetKey(e.x, e.y - 1)
	// Add Left
	neighbors[1] = GetKey(e.x - 1, e.y)
	// Add Right
	neighbors[2] = GetKey(e.x + 1, e.y)
	// Add Bottom
	neighbors[3] = GetKey(e.x, e.y + 1)
	return neighbors[:]
}

func PartA(input string) int {
	entries := GetEntries(input)
	sum := 0
	for _, entry := range entries {
		isLowest := true
		for _, neighbor := range entry.GetNeighbors() {
			if neighborEntry, ok := entries[neighbor]; ok {
				isLowest = isLowest && (entry.value < neighborEntry.value)
			}
		}
		if isLowest {
			sum += entry.value + 1
		}
	}
	return sum
}

func GetEntries(input string) map[int]*Entry {
	y := 0
	entries := map[int]*Entry{}
	common.SplitLines(input, func(s string) {
		for x, _ := range s {
			key := GetKey(x, y)
			entry := &Entry {
				value: common.CharAtoi(s[x]),
				x: x,
				y: y,
				explored: false,
			}
			entries[key] = entry
		}
		y += 1
	})
	return entries
}

func GetKey(x, y int) int {
	return y * 1000 + x
}

func PartB(input string) int {
	entries := GetEntries(input)
	h := NewMaxIntHeap()
	heap.Init(h)
	for _, entry := range entries {
		if !entry.explored && entry.value != 9 {
			size := ExploreBasin(entry, entries)
			heap.Push(h, size)
		}
	}
	return (*h)[0] * (*h)[1] * (*h)[2]
}

func ExploreBasin(entry *Entry, entries map[int]*Entry) int {
	queue := NewQueue()
	size := 0
	queue.Push(entry)
	for !queue.IsEmpty() {
		e := queue.Pop()
		if e.explored {
			continue
		}
		e.explored = true
		if e.value == 9 {
			continue
		}
		size += 1
		neighbors := e.GetNeighbors()
		for _, neighbor := range neighbors {
			if neighborEntry, ok := entries[neighbor]; ok && !neighborEntry.explored {
				queue.Push(neighborEntry)
			}
		}
	}
	return size
}