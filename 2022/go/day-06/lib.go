package lib

import (
	"strconv"
	"strings"
)

func PartA(input string) string {
	window := ""

	marker := 0
	for len(window) < 4 {
		if substr := string(input[marker]); strings.Contains(window, substr) {
			window = window[strings.Index(window, substr)+1:]
		}
		window += string(input[marker])
		marker++
	}

	return strconv.Itoa(marker)
}

func PartB(input string) string {
	window := ""

	marker := 0
	for len(window) < 14 {
		if substr := string(input[marker]); strings.Contains(window, substr) {
			window = window[strings.Index(window, substr)+1:]
		}
		window += string(input[marker])
		marker++
	}

	return strconv.Itoa(marker)
}
