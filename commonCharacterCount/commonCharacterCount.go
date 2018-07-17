package main

import (
	"fmt"
)

func main() {
	fmt.Println(commonCharacterCount("aabcc", "adcaa"))
}

func commonCharacterCount(s1 string, s2 string) int {
	commonCharacters := 0
	str1 := []byte(s1)
	str2 := []byte(s2)

	isCommon := func(char byte) bool {
		common := false

		for i, item := range str2 {
			if item == char {
				common = true
				str2 = append(str2[:i], str2[i+1:]...)
				break
			}
		}

		return common
	}

	for _, c := range str1 {
		if isCommon(c) {
			commonCharacters++
		}
	}

	return commonCharacters
}
