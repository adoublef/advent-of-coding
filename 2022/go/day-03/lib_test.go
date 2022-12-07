package day03

import (
	"testing"

	"github.com/hyphengolang/prelude/testing/is"
)

const input = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestPartA(t *testing.T) {
	t.Parallel()

	is := is.New(t)

	result := PartA(input)
	is.Equal(result, "157") // partA
}

func TestPartB(t *testing.T) {
	t.Parallel()

	is := is.New(t)

	result := PartB(input)
	is.Equal(result, "70") // partB
}
