package main

import (
	"fmt"
)

func main() {
	fmt.Println(allLongestStrings([]string{
		"aba",
		"aa",
		"ad",
		"vcd",
		"aba",
	}))
}

func allLongestStrings(inputArray []string) []string {
	targetLength := 0
	var longestStrings []string

	for _, str := range inputArray {
		switch {
		case len(str) == targetLength:
			targetLength = len(str)
			longestStrings = append(longestStrings, str)
		case len(str) > targetLength:
			targetLength = len(str)
			longestStrings = longestStrings[:0]
			longestStrings = append(longestStrings, str)
		}
	}

	return longestStrings
}
