package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(longestWord("Ready, steady, go!"))
}

func longestWord(text string) string {
	regex, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	clean := regex.ReplaceAllString(text, "")
	length := 0
	longestWord := ""

	for _, word := range strings.Split(clean, " ") {
		if len(word) > length {
			length = len(word)
			longestWord = word
		}
	}

	return longestWord
}
