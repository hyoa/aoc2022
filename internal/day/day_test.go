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
