package day

import (
	"hyoa/aoc2022/internal/utils"
	"strconv"
)

type Position struct {
	x, y int
}

type Grid struct {
	nodes map[Position]int
	sizeX int
	sizeY int
}

type Day8 struct {
	grid Grid
}

func (d *Day8) Init(path string) error {
	grid := Grid{nodes: make(map[Position]int)}
	x := 0
	maxY := 0
	for _, l := range utils.ReadTextFileLinesAsString(path) {
		y := 0
		for _, c := range l {
			v, _ := strconv.Atoi(string(c))
			grid.nodes[Position{x: x, y: y}] = v
			y++
		}

		if y > maxY {
			maxY = y
		}

		x++
	}

	d.grid = grid
	d.grid.sizeX = x
	d.grid.sizeY = maxY

	return nil
}

func (d *Day8) Step1() (Result, error) {
	visibles := 0

	moves := []Position{{y: -1}, {x: 1}, {y: 1}, {x: -1}}

	for p := range d.grid.nodes {
		if p.x == 0 || p.y == 0 || p.x+1 == d.grid.sizeX || p.y+1 == d.grid.sizeY {
			visibles++
			continue
		}

		for _, m := range moves {
			if walkIsVisible(d.grid, p, m, d.grid.nodes[p]) {
				visibles++
				break
			}
		}
	}

	return Result{Value: visibles}, nil
}

func (d *Day8) Step2() (Result, error) {
	moves := []Position{{y: -1}, {x: 1}, {y: 1}, {x: -1}}

	ttt := 0
	for p := range d.grid.nodes {
		tt := 1
		for _, m := range moves {
			tt *= walkQtyTreeView(d.grid, p, m, d.grid.nodes[p])
		}

		if tt > ttt {
			ttt = tt
		}
	}

	return Result{Value: ttt}, nil
}

func walkIsVisible(grid Grid, curr, direction Position, initialSize int) bool {
	next := Position{x: curr.x + direction.x, y: curr.y + direction.y}
	if grid.nodes[next] >= initialSize {
		return false
	}

	if _, ok := grid.nodes[next]; !ok {
		return true
	}

	return walkIsVisible(grid, next, direction, initialSize)
}

func walkQtyTreeView(grid Grid, curr, direction Position, initialSize int) int {
	next := Position{x: curr.x + direction.x, y: curr.y + direction.y}

	if _, ok := grid.nodes[next]; !ok {
		return 0
	}

	if grid.nodes[next] >= initialSize {
		return 1
	}

	return walkQtyTreeView(grid, next, direction, initialSize) + 1
}
