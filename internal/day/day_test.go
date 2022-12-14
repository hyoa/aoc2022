package day_test

import (
	"hyoa/aoc2022/internal/day"
	"testing"

	"github.com/joho/godotenv"
)

func TestDay(t *testing.T) {
	godotenv.Load("../../.env")

	type testCase struct {
		day           day.Day
		step1Expected interface{}
		step2Expected interface{}
	}

	cases := []testCase{
		{day: day.Day{Kind: day.InputKindExample, Runner: &day.Day1{}, Name: "day1"}, step1Expected: 24000, step2Expected: 45000},
		{day: day.Day{Kind: day.InputKindExample, Runner: &day.Day2{}, Name: "day2"}, step1Expected: 15, step2Expected: 12},
		{day: day.Day{Kind: day.InputKindExample, Runner: &day.Day3{}, Name: "day3"}, step1Expected: 157, step2Expected: 70},
		{day: day.Day{Kind: day.InputKindExample, Runner: &day.Day4{}, Name: "day4"}, step1Expected: 2, step2Expected: 4},
		{day: day.Day{Kind: day.InputKindExample, Runner: &day.Day5{}, Name: "day5"}, step1Expected: "CMZ", step2Expected: "MCD"},
		{day: day.Day{Kind: day.InputKindExample, Runner: &day.Day6{}, Name: "day6"}, step1Expected: 11, step2Expected: 26},
		{day: day.Day{Kind: day.InputKindExample, Runner: &day.Day7{}, Name: "day7"}, step1Expected: 95437, step2Expected: 24933642},
		{day: day.Day{Kind: day.InputKindExample, Runner: &day.Day8{}, Name: "day8"}, step1Expected: 21, step2Expected: 8},
		// {day: day.Day{Kind: day.InputKindExample, Runner: &day.Day9{}, Name: "day9"}, step1Expected: 13, step2Expected: 1},
		{day: day.Day{Kind: day.InputKindExample, Runner: &day.Day9{}, Name: "day9_2"}, step1Expected: 88, step2Expected: 36},
		{day: day.Day{Kind: day.InputKindExample, Runner: &day.Day10{}, Name: "day10"}, step1Expected: 13140, step2Expected: 0},
		{day: day.Day{Kind: day.InputKindExample, Runner: &day.Day13{}, Name: "day13"}, step1Expected: 13, step2Expected: 140},
		{day: day.Day{Kind: day.InputKindExample, Runner: &day.Day14{}, Name: "day14"}, step1Expected: 24, step2Expected: 93},
	}

	for _, c := range cases {
		t.Run(c.day.Name, func(t *testing.T) {
			c.day.Init()

			res1, err1 := c.day.Step1()

			if err1 != nil {
				t.Error(err1)
			}

			if c.step1Expected != res1.Value {
				t.Errorf("entry 1: expected %d got %d", c.step1Expected, res1.Value)
			}

			res2, err2 := c.day.Step2()

			if err2 != nil {
				t.Error(err2)
			}

			if c.step2Expected != res2.Value {
				t.Errorf("entry 2: expected %d got %d", c.step2Expected, res2.Value)
			}
		})
	}
}
