package main

import (
	_ "embed"
	"fmt"
	"strings"
)

// A/X => Rock
// B/Y => Paper
// C/Z => Scissors

//go:embed input.txt
var inputFile string

var scoresPerThing = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var counterPartToWin = map[string]string{
	"C": "X",
	"A": "Y",
	"B": "Z",
}

var counterPartToDraw = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var counterPartToLose = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}

func PartA(lines []string) any {
	score := 0

	for _, line := range lines {
		bits := strings.Split(line, " ")
		opponent := bits[0]
		me := bits[1]
		score += scoresPerThing[me]
		if me == counterPartToWin[opponent] {
			score += 6
		} else if me == counterPartToDraw[opponent] {
			score += 3
		}
	}

	return score
}

func main() {
	res := PartA(strings.Split(inputFile, "\n"))
	fmt.Println(res)
}
