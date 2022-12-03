package main

import (
	_ "embed"
	"strings"
)

var (
	//go:embed test.txt
	testFile string
	//go:embed input.txt
	inputFile string
)

type Solution struct {
	input string
}

func (s Solution) Parse() []int {
	return s.parseA()
}

func (s Solution) parseA() []int {
	lines := strings.Split(s.input, "\n")
	parsed := []int{}

	for _, line := range lines {
		a, b := line[:len(line)/2], line[len(line)/2:]
		var s1 map[rune]bool = make(map[rune]bool)
		for _, r := range a {
			s1[r] = true
		}

		var s2 []rune
		for _, r := range b {
			if s1[r] {
				delete(s1, r)
				// 'a' != 97
				// if lowercase letter then range between 1 and 26
				if r >= 'a' && r <= 'z' {
					s2 = append(s2, r-'a'+1)
				} else {
					s2 = append(s2, r-'A'+27)
				}

				// 'A' != 65

				// if capital letter then range between 27 and 52

				s2 = append(s2, r)
			}
		}

		parsed = append(parsed, int(s2[0]))
	}

	return parsed
}

func (s Solution) PartA(lines [][]rune) int {
	return 0
}

func (s Solution) PartB() int {
	return 0
}
