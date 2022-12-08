package lib

import (
	"strconv"
	"strings"
)

func PartA(input string) string {
	var res int
	for i := 0; i < len(input)-1; i++ {
		if window := input[i : i+4]; len(newDict(window).String()) == len(window) {
			return strconv.Itoa(i + len(window))
		}
	}

	return strconv.Itoa(res)
}

func PartB(input string) string {
	var res int
	for i := 0; i < len(input)-1; i++ {
		if window := input[i : i+14]; len(newDict(window).String()) == len(window) {
			return strconv.Itoa(i + len(window))
		}
	}

	return strconv.Itoa(res)
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

func (d Dict) Max() (rune, int) {
	var maxCh rune
	var maxI int
	for ch, i := range d {
		if d[ch] > d[maxCh] {
			maxCh = ch
			maxI = i
		}
	}

	return maxCh, maxI
}

/*
m
j
q
j
p
q
m
g
b
l
j
s
p
h
d
z
t
n
v
j
f
q
w
r
c
g
s
m
l
b
*/
