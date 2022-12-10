package day

import (
	"fmt"
	"hyoa/aoc2022/internal/utils"
	"strings"
)

type signal struct {
	kind  signalKind
	value int
}

type signalKind string

const (
	signalKindNoop signalKind = "noop"
	signalKindAdd  signalKind = "addx"
)

type Day10 struct {
	signals []signal
}

func (d *Day10) Init(path string) error {
	signals := make([]signal, 0)
	for _, l := range utils.ReadTextFileLinesAsString(path) {
		var kind string
		var value int

		fmt.Sscanf(l, "%s %d", &kind, &value)

		if kind == string(signalKindNoop) {
			signals = append(signals, signal{kind: signalKindNoop})
		} else {
			signals = append(signals, signal{kind: signalKindAdd, value: value})
		}
	}

	d.signals = signals
	return nil
}

func (d *Day10) Step1() (Result, error) {
	cycle := getCyclesValue(d.signals, 220)

	tt := 0
	k := []int{20, 60, 100, 140, 180, 220}
	for _, v := range k {
		tt += cycle[v-1] * v
	}

	return Result{Value: tt}, nil
}

func (d *Day10) Step2() (Result, error) {
	var crtDrawning [][]string

	var y, x int

	for k, v := range getCyclesValue(d.signals, 240) {
		if k%40 == 0 {
			x = 0
			crtDrawning = append(crtDrawning, []string{})
			y = len(crtDrawning) - 1
		}

		if x == v || x == v-1 || x == v+1 {
			crtDrawning[y] = append(crtDrawning[y], "#")
		} else {
			crtDrawning[y] = append(crtDrawning[y], ".")
		}

		x++
	}

	for _, v := range crtDrawning {
		fmt.Println(strings.Join(v, ""))
	}

	// Not testable
	return Result{Value: 0}, nil
}

func getCyclesValue(signals []signal, size int) []int {
	cycle := make([]int, size)
	cursor := 0

	cycle[0] = 1
	for i := 0; i < len(signals) && cursor < len(cycle)-1; i++ {
		cursor++

		cycle[cursor] = cycle[cursor-1]

		if signals[i].kind == signalKindAdd && cursor < len(cycle)-1 {
			cursor++
			cycle[cursor] = cycle[cursor-1] + signals[i].value
		}
	}

	return cycle
}
