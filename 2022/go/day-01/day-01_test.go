package day01

import (
	"strconv"
	"testing"

	"github.com/hyphengolang/prelude/testing/is"
)

const input = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestPartA(t *testing.T) {
	t.Parallel()

	is := is.New(t)

	result := PartA(input)
	is.Equal(result, "24000") // partA
}

func TestPartB(t *testing.T) {
	t.Parallel()

	is := is.New(t)

	result := PartB(input)
	is.Equal(result, strconv.Itoa(24000+11000+10000)) // partB
}
