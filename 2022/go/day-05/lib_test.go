package day05

import (
	"testing"

	"github.com/hyphengolang/prelude/testing/is"
)

const input = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestPartA(t *testing.T) {
	t.Parallel()

	is := is.New(t)

	result := PartA(input)
	is.Equal(result, "CMZ") // partA
}

func TestPartB(t *testing.T) {
	t.Parallel()

	is := is.New(t)

	result := PartB(input)
	is.Equal(result, "todo") // partB
}
