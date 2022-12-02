package rps

import (
	"errors"
)

// Key's are teh losing hand relative to the values which are the winning hand
//
//	map[loser]winner
var matcher = map[Hand]Hand{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

func LosesTo(h Hand) Hand {
	return matcher[h]
}

func WinsAgainst(h Hand) (Hand, error) {
	for w, l := range matcher {
		if h == l {
			return w, nil
		}
	}

	return -1, errors.New("invalid combination")
}

type Hand int

const (
	Rock Hand = iota + 1
	Paper
	Scissors
)

func (h Hand) Int() int { return int(h) }

func ParseHand(s string) (Hand, error) {
	switch s {
	case "A", "X":
		return Rock, nil
	case "B", "Y":
		return Paper, nil
	case "C", "Z":
		return Scissors, nil
	default:
		return -1, errors.New("invalid input")
	}
}

func (h Hand) Result(elf Hand) Result {
	switch {
	case win(h, elf):
		return Win
	case lose(h, elf):
		return Lose
	default:
		return Draw
	}
}

func win(a, b Hand) bool {
	return matcher[a] == b
}

func lose(a, b Hand) bool {
	return matcher[b] == a
}

type Result int

const (
	Lose Result = iota * 3
	Draw
	Win
)

func (r Result) Int() int { return int(r) }

func ParseResult(s string) (Result, error) {
	switch s {
	case "X":
		return Lose, nil
	case "Y":
		return Draw, nil
	case "Z":
		return Win, nil
	default:
		return -1, errors.New("invalid value")
	}
}

func (r Result) GetPair(h Hand) (Hand, error) {
	switch r {
	case Draw:
		return h, nil
	case Lose:
		return matcher[h], nil
	case Win:
		for w, l := range matcher {
			if l == h {
				return w, nil
			}
		}
		return -1, errors.New("invalid combination")
	default:
		return -1, errors.New("invalid combination")
	}
}
