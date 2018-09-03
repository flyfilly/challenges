package main

import (
	"fmt"
)

func main() {
	fmt.Println(alphabeticShift("crazy"))
}

func alphabeticShift(inputString string) string {
	newStr := ""
	swap := func(chr string) string {
		alphabet := "abcdefghijklmnopqrstuvwxyz"
		if chr == "z" {
			return "a"
		}

		for i := 0; i < len(alphabet); i++ {
			if chr == string(alphabet[i]) {
				return string(alphabet[i+1])
			}
		}

		return ""
	}

	for i := 0; i < len(inputString); i++ {
		newStr += swap(string(inputString[i]))
	}

	return newStr
}
