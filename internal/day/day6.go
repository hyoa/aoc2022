package day

import (
	"hyoa/aoc2022/internal/utils"
)

type Queue struct {
	node []string
	size int
}

type Day6 struct {
	input string
}

func (d *Day6) Init(path string) error {
	d.input = utils.ReadTextFile(path)
	return nil
}

func (d *Day6) Step1() (Result, error) {
	queue := Queue{size: 4}
	queue.init(d.input)

	idx := 0

	if queue.hasMarker() {
		return Result{Value: 4}, nil
	}

	for i := 4; i < len(d.input); i++ {
		queue.add(string(d.input[i]))
		if queue.hasMarker() {
			idx = i
			break
		}
	}

	return Result{Value: idx + 1}, nil
}

func (d *Day6) Step2() (Result, error) {
	queue := Queue{size: 14}
	queue.init(d.input)

	if queue.hasMarker() {
		return Result{Value: 14}, nil
	}

	idx := 0
	for i := 14; i < len(d.input); i++ {
		queue.add(string(d.input[i]))
		if queue.hasMarker() {
			idx = i
			break
		}
	}

	return Result{Value: idx + 1}, nil
}

func (q *Queue) init(s string) {
	q.node = make([]string, q.size)
	for i := 0; i < q.size; i++ {
		q.node[i] = string(s[i])
	}
}

func (q *Queue) add(s string) {
	for i := 0; i < q.size-1; i++ {
		q.node[i] = q.node[i+1]
	}
	q.node[q.size-1] = s
}

func (q *Queue) hasMarker() bool {
	chars := make(map[string]bool)
	for _, n := range q.node {
		if _, ok := chars[n]; ok {
			return false
		}

		chars[n] = true
	}

	return true
}
