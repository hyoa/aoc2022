package day

import (
	"fmt"
	"hyoa/aoc2022/internal/utils"
	"strings"
)

type Day14 struct {
	lowest           int
	occupiedPosition map[Position]string
}

func (d *Day14) Init(path string) error {
	lines := utils.ReadTextFileLinesAsString(path)

	rocksPosition := make(map[Position]string)
	var lowest int

	for _, l := range lines {
		d := strings.Split(l, " -> ")

		for i := 0; i < len(d)-1; i++ {

			var fromX, fromY, toX, toY int
			fmt.Sscanf(d[i], "%d,%d", &fromX, &fromY)
			fmt.Sscanf(d[i+1], "%d,%d", &toX, &toY)

			rocksPosition[Position{x: fromX, y: fromY}] = "#"
			rocksPosition[Position{x: toX, y: toY}] = "#"

			if fromY > lowest {
				lowest = fromY
			} else if toY > lowest {
				lowest = toY
			}

			if fromX == toX && fromY < toY {
				it := toY - fromY
				for j := 0; j < it; j++ {
					rocksPosition[Position{x: fromX, y: fromY + j}] = "#"
				}
			} else if fromX == toX && fromY > toY {
				it := fromY - toY
				for j := 0; j < it; j++ {
					rocksPosition[Position{x: toX, y: toY + j}] = "#"
				}
			} else if fromY == toY && fromX < toX {
				it := toX - fromX
				for j := 0; j < it; j++ {
					rocksPosition[Position{x: fromX + j, y: fromY}] = "#"
				}
			} else if fromY == toY && fromX > toX {
				it := fromX - toX
				for j := 0; j < it; j++ {
					rocksPosition[Position{x: toX + j, y: toY}] = "#"
				}
			}
		}
	}

	d.occupiedPosition = rocksPosition
	d.lowest = lowest
	return nil
}

const (
	MoveDown      string = "D"
	MoveDownLeft  string = "DL"
	MoveDownRight string = "DR"
)

func (d *Day14) Step1() (Result, error) {
	moves := map[string]Position{MoveDown: {y: 1}, MoveDownLeft: {y: 1, x: -1}, MoveDownRight: {y: 1, x: 1}}
	sandCount := 0

End:
	for {

		curr := Position{x: 500, y: 0}
		move := MoveDown
		for {
			next := Position{y: curr.y + moves[move].y, x: curr.x + moves[move].x}

			// if it == 23 {
			// 	fmt.Println(move)
			// 	printSandAndRock(d.occupiedPosition, curr)
			// }

			if _, ok := d.occupiedPosition[next]; ok {
				if move == MoveDown {
					move = MoveDownLeft
					continue
				} else if move == MoveDownLeft {
					move = MoveDownRight
					continue
				} else {
					d.occupiedPosition[curr] = "O"
					sandCount++
					break
				}
			} else {
				if next.y > d.lowest {
					break End
				}

				curr = next
				move = MoveDown
			}
		}

	}

	// fmt.Println(sandCount)
	// // printSandAndRock(d.occupiedPosition)
	return Result{Value: sandCount}, nil
}

func (d *Day14) Step2() (Result, error) {
	moves := map[string]Position{MoveDown: {y: 1}, MoveDownLeft: {y: 1, x: -1}, MoveDownRight: {y: 1, x: 1}}
	sandCount := 0

End:
	for {

		curr := Position{x: 500, y: 0}
		move := MoveDown
		for {
			next := Position{y: curr.y + moves[move].y, x: curr.x + moves[move].x}

			// if it == 23 {
			// 	fmt.Println(move)
			// 	printSandAndRock(d.occupiedPosition, curr)
			// }

			if next.y == d.lowest+2 {
				d.occupiedPosition[curr] = "O"
				sandCount++
				break
			} else if _, ok := d.occupiedPosition[next]; ok {
				if move == MoveDown {
					move = MoveDownLeft
					continue
				} else if move == MoveDownLeft {
					move = MoveDownRight
					continue
				} else {

					d.occupiedPosition[curr] = "O"
					sandCount++
					if curr.x == 500 && curr.y == 0 {
						break End
					}
					break
				}
			} else {
				curr = next
				move = MoveDown
			}
		}

	}

	p := 0
	for _, c := range d.occupiedPosition {
		if c == "O" {
			p++
		}
	}

	// printSandAndRock(d.occupiedPosition, Position{})
	return Result{Value: p}, nil
}

func printSandAndRock(node map[Position]string, track Position) {
	for y := 0; y < 20; y++ {
		line := make([]string, 0)

		for x := 480; x < 520; x++ {

			char := "."
			if v, ok := node[Position{x: x, y: y}]; ok {
				char = v
			}

			if x == track.x && y == track.y {
				char = "T"
			}

			line = append(line, char)
		}

		fmt.Println(strings.Join(line, ""))
	}
}
