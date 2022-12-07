package day03

import (
	"strconv"
	"strings"
)

func PartA(input string) string {
	rucksacks := strings.Split(input, "\n")

	var bs []Set
	for _, rucksack := range rucksacks {
		a, b := []rune(rucksack[:len(rucksack)/2]), []rune(rucksack[len(rucksack)/2:])
		bs = append(bs, newSet(a...).Union(b...))
	}

	var result int
	for _, set := range bs {
		for ch := range set {
			if ch < 97 {
				result += int(ch-'A') + 27
			} else {
				result += int(ch-'a') + 1
			}
		}
	}

	return strconv.Itoa(result)
}

func PartB(input string) string {
	rucksacks := strings.Split(input, "\n")

	var groups [][3]Set
	for i, rucksack := range rucksacks {
		j, s := i/3, []rune(rucksack)
		if i%3 == 0 {
			groups = append(groups, [3]Set{})
		}
		groups[j][i%3] = newSet(s...)
	}

	var res string
	for _, group := range groups {
		a, b, c := group[0], []rune(group[1].String()), []rune(group[2].String())
		res += a.Union(b...).Union(c...).String()
	}

	var result int
	for _, ch := range res {
		if ch < 97 {
			result += int(ch-'A') + 27
		} else {
			result += int(ch-'a') + 1
		}
	}

	return strconv.Itoa(result)
}

type Set map[rune]struct{}

func newSet(s ...rune) Set {
	m := map[rune]struct{}{}
	for _, ch := range s {
		m[ch] = struct{}{}
	}

	return m
}

func (s Set) Add(r rune) {
	s[r] = struct{}{}
}

func (s Set) Has(r rune) bool {
	_, ok := s[r]
	return ok
}

func (s Set) String() string {
	var sb strings.Builder
	for ch := range s {
		sb.WriteRune(ch)
	}
	return sb.String()
}

func (s Set) Union(b ...rune) Set {
	s1 := newSet()
	for _, ch := range b {
		if s.Has(ch) {
			s1.Add(ch)
		}
	}

	return s1
}
