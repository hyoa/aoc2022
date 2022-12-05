package day

import (
	"hyoa/aoc2022/internal/utils"
	"log"
	"strconv"
	"strings"
)

type Day4 struct {
	assignments []string
}

type section struct {
	start, end int
}

func (d *Day4) Init(path string) error {
	d.assignments = utils.ReadTextFileLinesAsString(path)
	return nil
}

func (d *Day4) Step1() (Result, error) {
	count := 0
	for _, assignment := range d.assignments {
		elves := strings.Split(assignment, ",")

		if overlap(getSection(elves[0]), getSection(elves[1])) {
			count++
		}
	}

	return Result{
		Kind:  ResultKindInt,
		Value: count,
	}, nil
}

func (d *Day4) Step2() (Result, error) {
	count := 0
	for _, assignment := range d.assignments {
		elves := strings.Split(assignment, ",")

		if overlapEvenPartially(getSection(elves[0]), getSection(elves[1])) {
			count++
		}
	}

	return Result{
		Kind:  ResultKindInt,
		Value: count,
	}, nil
}

func getSection(assignment string) section {
	sectionNumbers := strings.Split(assignment, "-")

	start, errStart := strconv.Atoi(sectionNumbers[0])
	end, errEnd := strconv.Atoi(sectionNumbers[1])

	if errStart != nil || errEnd != nil {
		log.Fatalln("error on atoi")
	}

	return section{
		start: start,
		end:   end,
	}
}

func overlap(section1, section2 section) bool {
	if section1.start <= section2.start && section1.end >= section2.end {
		return true
	}

	if section2.start <= section1.start && section2.end >= section1.end {
		return true
	}

	return false
}

func overlapEvenPartially(section1, section2 section) bool {
	if overlap(section1, section2) {
		return true
	}

	if section2.start >= section1.start && section2.start <= section1.end {
		return true
	}

	if section1.start >= section2.start && section1.start <= section2.end {
		return true
	}

	return false
}
