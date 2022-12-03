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
	fmt.Println(days[index].Step1())
	fmt.Println(days[index].Step2())
}
