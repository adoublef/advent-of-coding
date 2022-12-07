package day04

import (
	"testing"

	"github.com/hyphengolang/prelude/testing/is"
)

const input = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestPartA(t *testing.T) {
	t.Parallel()

	is := is.New(t)

	result := PartA(input)
	is.Equal(result, "2") // partA
}

func TestPartB(t *testing.T) {
	t.Parallel()

	is := is.New(t)

	result := PartB(input)
	is.Equal(result, "4") // partB
}
