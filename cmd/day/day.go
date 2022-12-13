package main

import (
	"fmt"
	"hyoa/aoc2022/internal/day"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	input := os.Args[2]

	kind := day.InputKindReal
	if input == "example" {
		kind = day.InputKindExample
	}

	days := []day.Day{
		{Runner: &day.Day1{}, Name: "day1", Kind: kind},
		{Runner: &day.Day2{}, Name: "day2", Kind: kind},
		{Runner: &day.Day3{}, Name: "day3", Kind: kind},
		{Runner: &day.Day4{}, Name: "day4", Kind: kind},
		{Runner: &day.Day5{}, Name: "day5", Kind: kind},
		{Runner: &day.Day6{}, Name: "day6", Kind: kind},
		{Runner: &day.Day7{}, Name: "day7", Kind: kind},
		{Runner: &day.Day8{}, Name: "day8", Kind: kind},
		{Runner: &day.Day9{}, Name: "day9", Kind: kind},
		{Runner: &day.Day10{}, Name: "day10", Kind: kind},
		{Runner: &day.Day10{}, Name: "day10", Kind: kind},
		{Runner: &day.Day10{}, Name: "day10", Kind: kind},
		{Runner: &day.Day13{}, Name: "day13", Kind: kind},
	}

	n, errAtoi := strconv.Atoi(os.Args[1])
	index := n - 1

	if errAtoi != nil {
		log.Fatalln(errAtoi)
	}

	if len(days) <= index {
		log.Fatalln("day not found: ", n)
	}

	fmt.Printf("run day %d with %s inputs \r\n", n, kind)
	days[index].Init()

	step1, _ := days[index].Step1()
	step2, _ := days[index].Step2()

	fmt.Println(step1.Value)
	fmt.Println(step2.Value)
}
