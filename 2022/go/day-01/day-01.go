package day01

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func PartA(input string) string {
	supplies := parse(input)

	var max = 0
	for _, supply := range supplies {
		max = int(math.Max(float64(supply), float64(max)))
	}

	return strconv.Itoa(max)
}

func PartB(input string) string {
	supplies := parse(input)

	sort.Ints(supplies)

	var sum = 0
	for _, n := range supplies[len(supplies)-3:] {
		sum += n
	}

	return strconv.Itoa(sum)
}

func parse(input string) []int {
	supplies := strings.Split(input, "\n\n")

	var sums []int
	for _, supply := range supplies {
		calories := strings.Split(supply, "\n")

		var sum = 0
		for _, calorie := range calories {
			n, _ := strconv.Atoi(calorie)
			sum += n
		}

		sums = append(sums, sum)
	}

	return sums
}
