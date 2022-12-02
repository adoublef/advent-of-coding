// --- Day 2: Rock Paper Scissors ---
//
// The Elves begin to set up camp on the beach. To decide whose tent gets to be closest to the snack storage, a giant Rock Paper Scissors tournament is already in progress.
//
// Rock Paper Scissors is a game between two players. Each game contains many rounds; in each round, the players each simultaneously choose one of Rock, Paper, or Scissors using a hand shape. Then, a winner for that round is selected: Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock. If both players choose the same shape, the round instead ends in a draw.
//
// Appreciative of your help yesterday, one Elf gives you an encrypted strategy guide (your puzzle input) that they say will be sure to help you win. "The first column is what your opponent is going to play: A for Rock, B for Paper, and C for Scissors. The second column--" Suddenly, the Elf is called away to help with someone's tent.
//
// The second column, you reason, must be what you should play in response: X for Rock, Y for Paper, and Z for Scissors. Winning every time would be suspicious, so the responses must have been carefully chosen.
//
// The winner of the whole tournament is the player with the highest score. Your total score is the sum of your scores for each round. The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).
package main

import (
	_ "embed"
	"fmt"
	"strings"

	"aoc.day-2/rps"
)

//go:embed test.txt
var testTxt string

//go:embed input.txt
var inputTxt string

func main() {
	fmt.Printf("result of part_a is %d\n", partA(strings.Split(inputTxt, "\n")))
	fmt.Printf("result of part_b is %d\n", partB(strings.Split(inputTxt, "\n")))
}

func partA(lines []string) int {
	var n = 0
	for _, line := range lines {
		inputs := strings.Split(line, " ")
		elf, _ := rps.ParseHand(inputs[0])
		me, _ := rps.ParseHand(inputs[1])

		n += me.Int() + me.Result(elf).Int()
	}

	return n
}

func partB(lines []string) int {
	var n = 0
	for _, line := range lines {
		inputs := strings.Split(line, " ")
		elf, _ := rps.ParseHand(inputs[0])
		r, _ := rps.ParseResult(inputs[1])

		me, _ := r.GetPair(elf)
		n += r.Int() + me.Int()
	}

	return n
}
