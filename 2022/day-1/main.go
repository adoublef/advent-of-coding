package main

import (
	"bufio"
	_ "embed"
	"os"
	"sort"
	"strconv"
)

var (
	//go:embed test.txt
	test string
	//go:embed input.txt
	input string
)

func main() {
	input := readInput("./2022/day-1/test.txt")
	println(partB(input))
}

func partA(input []int) int {
	var max = 0
	for _, num := range input {
		if num > max {
			max = num
		}
	}

	return max
}

func partB(input []int) int {
	sort.Slice(input, func(i, j int) bool { return i > j })

	return input[0] + input[1] + input[2]
}

func readInput(path string) []int {
	file, _ := os.Open(path)
	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)

	var input []int
	var acc = 0
	for sc.Scan() {
		line := sc.Text()
		if num, err := strconv.Atoi(line); err != nil {
			input = append(input, acc)
			acc = 0
		} else {
			acc += num
		}
	}

	input = append(input, acc)

	return input
}
