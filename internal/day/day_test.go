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
		step1Expected int
		step2Expected int
	}

	cases := []testCase{
		{day: day.Day{Kind: day.InputKindExample, Runner: &day.Day1{}, Name: "day1"}, step1Expected: 24000, step2Expected: 45000},
	}

	for _, c := range cases {
		t.Run(c.day.Name, func(t *testing.T) {
			c.day.Init()

			res1, err1 := c.day.Step1()

			if err1 != nil {
				t.Error(err1)
			}

			if c.step1Expected != res1 {
				t.Errorf("entry 1: expected %d got %d", c.step1Expected, res1)
			}

			res2, err2 := c.day.Step2()

			if err2 != nil {
				t.Error(err2)
			}

			if c.step2Expected != res2 {
				t.Errorf("entry 1: expected %d got %d", c.step2Expected, res2)
			}
		})
	}
}
