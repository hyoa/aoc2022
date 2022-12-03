package day

import (
	"hyoa/aoc2022/internal/utils"
	"strings"
)

type Day3 struct {
	data []string
}

func (d *Day3) Init(path string) error {
	d.data = utils.ReadTextFileLinesAsString(path)
	return nil
}

func (d *Day3) Step1() (int, error) {
	tt := 0
	for _, content := range d.data {
		cSplit := strings.Split(content, "")
		middle := len(cSplit) / 2
		intersect := intersectSlice(cSplit[:middle], cSplit[middle:])

		for _, i := range intersect {
			tt += getUnicodeValue(i)
		}
	}

	return tt, nil
}

func (d *Day3) Step2() (int, error) {
	tt := 0

	for i := 0; i < len(d.data); i += 3 {
		intersect := intersectSlice(strings.Split(d.data[i], ""), strings.Split(d.data[i+1], ""))
		intersectAll := intersectSlice(intersect, strings.Split(d.data[i+2], ""))

		for _, i := range intersectAll {
			tt += getUnicodeValue(i)
		}
	}

	return tt, nil
}

func intersectSlice(s1, s2 []string) []string {
	s1MappedValue := make(map[string]bool)

	for _, it := range s1 {
		s1MappedValue[it] = false
	}

	for _, it := range s2 {
		if _, ok := s1MappedValue[it]; ok {
			s1MappedValue[it] = true
		}
	}

	var intersect []string

	for k, b := range s1MappedValue {
		if b {
			intersect = append(intersect, k)
		}
	}

	return intersect
}

func getUnicodeValue(char string) int {
	v := int([]rune(char)[0])

	if v >= 97 {
		return v - 96
	}

	return v - 38
}
