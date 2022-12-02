package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hyphengolang/prelude/testing/is"
)

// lost = 0
// draw = 3
// win = 6

// a <-> rock :: paper means win 2
// b <-> paper :: scissor means win 3
// c <-> scissor :: rock means win 1

// y <-> paper
// x <-> rock
// z <-> scissors

func TestPartA(t *testing.T) {
	is := is.New(t)

	type testcase struct {
		input  string
		result int
	}

	tt := []testcase{
		{input: testTxt, result: 15},
	}

	for i, tc := range tt {
		name := fmt.Sprintf("test_%d", i+1)
		t.Run(name, func(t *testing.T) {
			lines := strings.Split(tc.input, "\n")

			result := partA(lines)

			is.Equal(result, tc.result) // correct result
		})
	}
}

func TestPartB(t *testing.T) {
	is := is.New(t)

	type testcase struct {
		input  string
		result int
	}

	tt := []testcase{
		{input: testTxt, result: 12},
	}

	for i, tc := range tt {
		name := fmt.Sprintf("test_%d", i+1)
		t.Run(name, func(t *testing.T) {
			lines := strings.Split(tc.input, "\n")

			result := partB(lines)

			is.Equal(result, tc.result) // correct result
		})
	}
}
