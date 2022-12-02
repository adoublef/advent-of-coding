package main

import (
	"testing"

	"github.com/hyphengolang/prelude/testing/is"
)

func TestSolution(t *testing.T) {
	is := is.New(t)

	input := readInput("test.txt")
	is.Equal(24_000, partA(input)) // ok
}

func TestReadInput(t *testing.T) {
	is := is.New(t)

	input := readInput("test.txt")

	is.Equal(6000, input[0])             // first value == 6000
	is.Equal(10000, input[len(input)-1]) // last value == 10000
	is.Equal(5, len(input))              // array length is 5
}
