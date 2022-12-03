package main

import (
	"fmt"
	"testing"

	"github.com/hyphengolang/prelude/testing/is"
)

func TestPartA(t *testing.T) {
	is := is.New(t)

	// sol := Solution{testFile}
	sol := Solution{inputFile}

	lines := sol.Parse()

	// sum int slice
	sum := 0
	for _, v := range lines {
		sum += v
	}

	is.Equal(sum, 0) // sum != 0

	// is.Equal(len(lines), 6) // test.txt has 6 lines
	// is.Equal('A', 102)                                      // 'A' != 65
	// is.Equal('a', 102)                                      // 'a' != 97
	// is.Equal(lines, []string{"p", "L", "P", "v", "t", "s"}) //test.txt
}

func TestPartA1(t *testing.T) {
	is := is.New(t)

	type testcase struct {
		input  string
		output string
	}

	testcases := []testcase{
		{input: "vJrwpWtwJgWrhcsFMMfFFhFp", output: "p"},
		{input: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", output: "L"},
		{input: "PmmdzqPrVvPwwTWBwg", output: "P"},
		{input: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", output: "v"},
		{input: "ttgJtRGJQctTZtZT", output: "t"},
		{input: "CrZsJsPPZsGzwwsLwLmpwMDw", output: "s"},
	}

	for i, tc := range testcases {
		name := fmt.Sprintf("testcase %d", i)
		t.Run(name, func(t *testing.T) {
			a, b := tc.input[:len(tc.input)/2], tc.input[len(tc.input)/2:]

			var s1 map[rune]bool = make(map[rune]bool)
			for _, r := range a {
				s1[r] = true
			}

			var s2 []rune
			for _, r := range b {
				if s1[r] {
					delete(s1, r)
					s2 = append(s2, r)
				}
			}

			is.Equal(string(s2), tc.output) // output != tc.output
		})
	}
}

//

const pb = `
s t V W Z z c Q L B P l P B s P W t Q h T J c z t b L l d G S q b t C Q G M z S M s L n d H L r N m q R C g L b n n S Z l q d g c r Q M C p r G M J s h w N m J b h W S D j v l m S p Q H t d N G r P J B l p z j l D j n L W J Q R n b h p S d N j V g t v c h b v D t m v n J m W R F C R t D W Z r h r C s p l S n g P j m J r w f M c R H P v b D d V N p N b B C b R l H j W T Q h T v Q b h V D b n m t g S l b d w W G J w B l V m W M Q p g V H c F l C r B s P B V h C L Z F P R q n b V Q Z q w l j d Q r m p J F G P V t t Z p N d R h r m B g S d F R m q T V L s H w p S l H C L n t D L C l H M r d z v P
`
