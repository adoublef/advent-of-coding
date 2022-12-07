package day02

import (
	"strconv"
	"strings"
)

type Hand int

func newHand(s string) Hand {
	switch s {
	case "A", "X":
		return HandRock
	case "B", "Y":
		return HandPaper
	case "C", "Z":
		return HandScissors
	}

	// if error return `-1`
	return -1
}

func (h Hand) Int() int { return int(h) }

func (h Hand) Result(b Hand) Game {

	switch {
	case h == b:
		return GameDraw
	case matcher[h] == b:
		return GameWin
	case matcher[b] == h:
		return GameLose
	}

	// if error return `-1`
	return -1
}

const (
	HandRock Hand = iota + 1
	HandPaper
	HandScissors
)

type Game int

func (g Game) Int() int { return int(g) }

func newGame(s string) Game {
	switch s {
	case "X":
		return GameLose
	case "Y":
		return GameDraw
	case "Z":
		return GameWin
	}

	// if error return `-1`
	return -1
}

func (g Game) Eval(h Hand) int {

	var v Hand
	switch g {
	case GameLose:
		v = matcher[h]
	case GameDraw:
		v = h
	case GameWin:
		for w, l := range matcher {
			if l == h {
				v = w
			}
		}
	}

	return g.Int() + v.Int()
}

const (
	GameLose Game = iota * 3
	GameDraw
	GameWin
)

var matcher = map[Hand]Hand{
	HandRock:     HandScissors,
	HandPaper:    HandRock,
	HandScissors: HandPaper,
}

func PartA(input string) string {
	games := strings.Split(input, "\n")

	var results int
	for _, game := range games {
		hands := strings.Split(game, " ")
		elf, me := newHand(hands[0]), newHand(hands[1])

		results += me.Int() + me.Result(elf).Int()
	}

	return strconv.Itoa(results)
}

func PartB(input string) string {
	games := strings.Split(input, "\n")

	var results int
	for _, game := range games {
		hands := strings.Split(game, " ")
		elf, result := newHand(hands[0]), newGame(hands[1])

		results += result.Eval(elf)
	}

	return strconv.Itoa(results)
}
