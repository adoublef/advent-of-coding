package lib

import (
	"strconv"
	"strings"
)

func PartA(input string) string {
	window := ""

	marker := 0
	for len(window) < 4 {
		s := string(input[marker])
		// if strings.Contains(window, s) {
		window = window[strings.Index(window, s)+1:]
		// }
		window += string(input[marker])
		marker++
	}

	return strconv.Itoa(marker)
}

func PartB(input string) string {
	window := ""

	marker := 0
	for len(window) < 14 {
		if s := string(input[marker]); strings.Contains(window, s) {
			window = window[strings.Index(window, s)+1:]
		}
		window += string(input[marker])
		marker++
	}

	return strconv.Itoa(marker)
}
