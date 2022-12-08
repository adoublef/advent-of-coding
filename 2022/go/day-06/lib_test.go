package lib

import (
	"fmt"
	"testing"

	"github.com/hyphengolang/prelude/testing/is"
)

func TestPartA(t *testing.T) {
	is := is.New(t)

	type testcase struct {
		input  string
		marker string
	}

	tt := []testcase{
		{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", marker: "7"},
		{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", marker: "5"},
		{input: "nppdvjthqldpwncqszvftbrmjlhg", marker: "6"},
		{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", marker: "10"},
		{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", marker: "11"},
	}

	for i, tc := range tt {
		name := fmt.Sprintf("test_%d", i)
		t.Run(name, func(t *testing.T) {
			result := PartA(tc.input)
			is.Equal(result, tc.marker) // ok
		})
	}
}

func TestPartB(t *testing.T) {
	is := is.New(t)

	type testcase struct {
		input  string
		marker string
	}

	tt := []testcase{
		{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", marker: "19"},
		{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", marker: "23"},
		{input: "nppdvjthqldpwncqszvftbrmjlhg", marker: "23"},
		{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", marker: "29"},
		{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", marker: "26"},
	}

	for i, tc := range tt {
		name := fmt.Sprintf("test_%d", i+1)
		t.Run(name, func(t *testing.T) {
			result := PartB(tc.input)
			is.Equal(result, tc.marker) // ok
		})
	}
}
