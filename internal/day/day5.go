package day

import (
	"fmt"
	"hyoa/aoc2022/internal/utils"
	"regexp"
	"strings"
)

type Day5 struct {
	moves  []Move
	stacks []Stack
	path   string
}

type Move struct {
	qty, start, end int
}

type Stack struct {
	head   *StackItem
	length int
}

type StackItem struct {
	value string
	prev  *StackItem
}

func (d *Day5) Init(path string) error {
	inputs := utils.ReadTextFileLinesAsString(path)
	var splitIndex int
	for k, i := range inputs {
		if strings.Contains(i, "move") {
			splitIndex = k
			break
		}
	}

	stacks := make([]Stack, 0)

	stacksContent := inputs[:splitIndex-2]

	rValue, _ := regexp.Compile("[A-Z]")
	for i := len(stacksContent); i != 0; i-- {
		stackIdx := 0
		for k, c := range stacksContent[i-1] {
			if k%4 == 0 {
				stackIdx++
				continue
			}

			if len(stacks) < stackIdx {
				stacks = append(stacks, Stack{})
			}

			if rValue.MatchString(string(c)) {
				stacks[stackIdx-1].push(StackItem{value: string(c)})
			}
		}
	}
	d.stacks = stacks

	moves := make([]Move, 0)
	for _, move := range inputs[splitIndex:] {
		var qty, start, end int
		fmt.Sscanf(move, "move %d from %d to %d", &qty, &start, &end)
		moves = append(moves, Move{qty: qty, start: start, end: end})
	}
	d.moves = moves

	d.path = path
	return nil
}

func (d *Day5) Step1() (Result, error) {
	for _, move := range d.moves {
		for i := 0; i < move.qty; i++ {
			it := d.stacks[move.start-1].pop()
			d.stacks[move.end-1].push(it)
		}
	}

	crates := make([]string, 0)
	for _, s := range d.stacks {
		crates = append(crates, s.pop().value)
	}

	return Result{Kind: ResultKindString, Value: strings.Join(crates, "")}, nil
}

func (d *Day5) Step2() (Result, error) {
	d.Init(d.path)

	// TODO remove one for by splitting 2 stack and reattach them correctly
	for _, move := range d.moves {
		cratesToMove := make([]StackItem, 0)
		for i := 0; i < move.qty; i++ {
			it := d.stacks[move.start-1].pop()
			cratesToMove = append(cratesToMove, it)
		}

		for i := len(cratesToMove); i != 0; i-- {
			d.stacks[move.end-1].push(cratesToMove[i-1])
		}
	}

	crates := make([]string, 0)
	for _, s := range d.stacks {
		crates = append(crates, s.pop().value)
	}

	return Result{Kind: ResultKindString, Value: strings.Join(crates, "")}, nil
}

func (s *Stack) push(item StackItem) {
	s.length++

	if s.head == (&StackItem{}) {
		s.head = &item
		return
	}

	item.prev = s.head
	s.head = &item
}

func (s *Stack) pop() StackItem {
	if s.length == 0 {
		return StackItem{}
	}

	s.length--

	item := s.head

	if s.length <= 0 {
		s.length = 0
		item.prev = &StackItem{}
		s.head = &StackItem{}

		return *item
	}

	s.head = item.prev
	item.prev = &StackItem{}

	return *item
}
