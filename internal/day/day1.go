package day

import (
	"hyoa/aoc2022/internal/utils"
	"sort"
	"strconv"
)

type Day1 struct {
	data []string
}

func (d *Day1) Init(path string) error {
	d.data = utils.ReadTextFileLinesAsString(path)
	return nil
}

func (d *Day1) Step1() (Result, error) {
	caloriesByElves := getTotalCaloriesByElves(d.data)

	return Result{
		Kind:  ResultKindInt,
		Value: caloriesByElves[len(caloriesByElves)-1],
	}, nil
}

func (d *Day1) Step2() (Result, error) {
	caloriesByElves := getTotalCaloriesByElves(d.data)

	tt := 0
	for _, c := range caloriesByElves[len(caloriesByElves)-3:] {
		tt += c
	}

	return Result{
		Kind:  ResultKindInt,
		Value: tt,
	}, nil
}

func getTotalCaloriesByElves(data []string) []int {
	caloryByElves := make([]int, 0)
	index := 0

	for _, d := range data {
		if d == "" {
			index++
			continue
		}

		if len(caloryByElves) < index+1 {
			caloryByElves = append(caloryByElves, 0)
		}

		v, _ := strconv.Atoi(d)
		caloryByElves[index] += v
	}

	sort.Slice(caloryByElves, func(i, j int) bool {
		return caloryByElves[i] < caloryByElves[j]
	})

	return caloryByElves
}
