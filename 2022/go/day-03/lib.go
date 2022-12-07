package day03

import (
	"strconv"
	"strings"
)

func PartA(input string) string {
	rucksacks := strings.Split(input, "\n")

	var ds []Dict
	for _, rucksack := range rucksacks {
		a, b := rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]
		ds = append(ds, newDict(a).Union(b))
	}

	var result int
	for _, d := range ds {
		for ch := range d {
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

	var groups []string
	for i, rucksack := range rucksacks {
		j, s := i/3, rucksack
		if i%3 == 0 {
			groups = append(groups, "")
		}
		groups[j] += newDict(s).String()
	}

	var res string
	for _, group := range groups {
		res += string(newDict(group).Max())
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

type Dict map[rune]int

func newDict(s string) Dict {
	m := map[rune]int{}
	for _, ch := range s {
		m[ch]++
	}

	return m
}

func (d Dict) Add(r rune) {
	d[r]++
}

func (d Dict) Has(r rune) bool {
	_, ok := d[r]
	return ok
}

func (d Dict) String() string {
	var sb strings.Builder
	for ch := range d {
		sb.WriteRune(ch)
	}
	return sb.String()
}

func (d Dict) Union(b string) Dict {
	s1 := newDict("")
	for _, ch := range b {
		if d.Has(ch) {
			s1.Add(ch)
		}
	}

	return s1
}

func (d Dict) Max() rune {
	var max rune
	for ch := range d {
		if d[ch] > d[max] {
			max = ch
		}
	}

	return max
}
