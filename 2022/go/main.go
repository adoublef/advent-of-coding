package main

import (
	d "day-06"
	"embed"
	"fmt"
)

//go:embed assets/*
var assets embed.FS

func main() {
	input, _ := assets.ReadFile("assets/day-06.txt")

	r := d.PartA(string(input))
	fmt.Printf("result of part A is := %s\n", r)

	r = d.PartB(string(input))
	fmt.Printf("result of part B is := %s\n", r)
}
