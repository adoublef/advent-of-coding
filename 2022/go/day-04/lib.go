package day04

import (
	"strconv"
	"strings"
)

func PartA(input string) string {
	ranges := parse(input)

	var res int
	for _, r := range ranges {
		if r[0].Contains(r[1]) || r[1].Contains(r[0]) {
			res++
		}
	}

	return strconv.Itoa(res)
}

func PartB(input string) string {
	ranges := parse(input)

	var res int
	for _, r := range ranges {
		if r[0].Overlaps(r[1]) || r[1].Overlaps(r[0]) {
			res++
		}
	}

	return strconv.Itoa(res)
}

type Range [2]int

func newRange(s string) Range {
	r := strings.Split(s, "-")
	a, _ := strconv.Atoi(r[0])
	b, _ := strconv.Atoi(r[1])
	return Range{a, b}
}

func (r Range) Overlaps(other Range) bool {
	return (r[0] <= other[0] && other[0] <= r[1]) || (r[0] <= other[1] && other[1] <= r[1])
}

func (r Range) Contains(other Range) bool {
	return r[0] <= other[0] && other[1] <= r[1]
}

func parse(input string) [][2]Range {
	lines := strings.Split(input, "\n")

	var ranges [][2]Range
	for _, line := range lines {
		v := strings.Split(line, ",")

		a, b := newRange(v[0]), newRange(v[1])
		ranges = append(ranges, [2]Range{a, b})
	}

	return ranges
}
