package day02

import (
	"testing"

	"github.com/hyphengolang/prelude/testing/is"
)

const input = `A Y
B X
C Z`

func TestPartA(t *testing.T) {
	t.Parallel()

	is := is.New(t)

	result := PartA(input)
	is.Equal(result, "15") // partA
}

func TestPartB(t *testing.T) {
	t.Parallel()

	is := is.New(t)

	result := PartB(input)
	is.Equal(result, "12") // partB
}
