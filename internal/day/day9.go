package day

import (
	"fmt"
	"hyoa/aoc2022/internal/utils"
	"strings"
)

type D9Position struct {
	x, y int
}

var D9MoveDirection = map[string]D9Position{"U": {y: -1}, "R": {x: 1}, "D": {y: 1}, "L": {x: -1}, "UR": {x: 1, y: -1}, "DR": {x: 1, y: 1}, "DL": {x: -1, y: 1}, "UL": {x: -1, y: -1}}
var MoveAvailable = []D9Position{{y: -1}, {x: 1}, {y: 1}, {x: -1}, {x: 1, y: -1}, {x: 1, y: 1}, {x: -1, y: 1}, {x: -1, y: -1}}

type D9Move struct {
	direction D9Position
	count     int
	name      string
}

type Day9 struct {
	moves []D9Move
}

func (d *Day9) Init(path string) error {
	moves := make([]D9Move, 0)
	for _, l := range utils.ReadTextFileLinesAsString(path) {
		var dir string
		var count int

		fmt.Sscanf(l, "%s %d", &dir, &count)
		moves = append(moves, D9Move{
			direction: D9MoveDirection[dir],
			count:     count,
			name:      dir,
		})
	}

	d.moves = moves
	return nil
}

type rope struct {
	knots []knot
}

type knot struct {
	curr D9Position
	prev D9Position
}

func (d *Day9) Step1() (Result, error) {
	rope := make([]knot, 2)
	for i := 0; i < len(rope); i++ {
		rope[i] = knot{curr: D9Position{x: 0, y: 0}, prev: D9Position{x: 0, y: 0}}
	}

	seen := make(map[D9Position]bool)

	seen[rope[len(rope)-1].curr] = true
	tt := 0
	for _, m := range d.moves {
		for i := 0; i < m.count; i++ {
			tt++

			for idx := range rope {
				if idx == 0 {
					rope[idx].prev = rope[idx].curr
					rope[idx].curr = D9Position{x: rope[idx].curr.x + m.direction.x, y: rope[idx].curr.y + m.direction.y}

					continue
				}

				if isAdjacent(rope[idx], rope[idx-1]) {
					continue
				}

				rope[idx].prev = rope[idx].curr
				rope[idx].curr = rope[idx-1].prev

				if idx == len(rope)-1 {
					seen[rope[idx].curr] = true
				}
			}
		}
	}

	return Result{Value: len(seen)}, nil
}

func (d *Day9) Step2() (Result, error) {
	rope := make([]knot, 10)
	for i := 0; i < len(rope); i++ {
		rope[i] = knot{curr: D9Position{x: 15, y: 15}, prev: D9Position{x: 15, y: 15}}
	}

	seen := make(map[D9Position]bool)

	seen[rope[len(rope)-1].curr] = true
	tt := 0
	for _, m := range d.moves {

		for i := 0; i < m.count; i++ {

			for idx := range rope {
				if idx == 0 {
					rope[idx].prev = rope[idx].curr
					rope[idx].curr = D9Position{x: rope[idx].curr.x + m.direction.x, y: rope[idx].curr.y + m.direction.y}
					continue
				}

				if isAdjacent(rope[idx], rope[idx-1]) {

					continue
				}

				rope[idx].prev = rope[idx].curr
				rope[idx].curr = newTailPosition(rope[idx], rope[idx-1])

				if idx == len(rope)-1 {
					seen[rope[idx].curr] = true
				}
			}
		}
		tt++

	}

	return Result{Value: len(seen)}, nil
}

func isAdjacent(tail, head knot) bool {
	if tail.curr.x == head.curr.x && tail.curr.y == head.curr.y {
		return true
	}

	for _, m := range MoveAvailable {
		p := D9Position{x: tail.curr.x + m.x, y: tail.curr.y + m.y}
		if p == head.curr {
			return true
		}
	}

	return false
}

func newTailPosition(tail, head knot) D9Position {
	if head.curr.x == tail.curr.x && head.curr.y < tail.curr.y {
		return D9Position{x: tail.curr.x, y: tail.curr.y - 1}
	} else if head.curr.x == tail.curr.x && head.curr.y > tail.curr.y {
		return D9Position{x: tail.curr.x, y: tail.curr.y + 1}
	} else if head.curr.y == tail.curr.y && head.curr.x < tail.curr.x {
		return D9Position{x: tail.curr.x - 1, y: tail.curr.y}
	} else if head.curr.y == tail.curr.y && head.curr.x > tail.curr.x {
		return D9Position{x: tail.curr.x + 1, y: tail.curr.y}
	} else if head.curr.x > tail.curr.x && head.curr.y < tail.curr.y {
		return D9Position{x: tail.curr.x + 1, y: tail.curr.y - 1}
	} else if head.curr.x > tail.curr.x && head.curr.y > tail.curr.y {
		return D9Position{x: tail.curr.x + 1, y: tail.curr.y + 1}
	} else if head.curr.x < tail.curr.x && head.curr.y < tail.curr.y {
		return D9Position{x: tail.curr.x - 1, y: tail.curr.y - 1}
	} else {
		return D9Position{x: tail.curr.x - 1, y: tail.curr.y + 1}
	}
}

func printRope(x, y int, rope []knot) {
	grid := make([][]string, 0)
	for i := 0; i < x; i++ {
		grid = append(grid, make([]string, 0))
		for j := 0; j < y; j++ {
			grid[i] = append(grid[i], ".")
		}
	}

	for k, p := range rope {
		grid[p.curr.y][p.curr.x] = fmt.Sprintf("%d", k)
	}

	grid[15][15] = "s"

	fmt.Println("--------")
	for _, g := range grid {
		fmt.Println(strings.Join(g, ""))
	}
}
