package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(longestDigitsPrefix("0123456789"))
}

func longestDigitsPrefix(inputString string) string {
	idx := 0

	for i, char := range inputString {
		idx = i
		fmt.Println(string(char), unicode.IsDigit(char))
		if !unicode.IsDigit(char) {
			break
		}
	}

	if (idx + 1) == len(inputString) {
		return inputString
	} else {
		return inputString[0:idx]
	}
}
