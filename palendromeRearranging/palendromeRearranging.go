package main

import "fmt"

func main() {
	fmt.Println(palindromeRearranging("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaabc"))
}

func palindromeRearranging(inputString string) bool {
	parts, odds := make(map[rune]int), 0

	for _, char := range inputString {
		parts[char]++
	}

	for part := range parts {
		if parts[part]%2 != 0 {
			odds++
		}
	}

	return odds <= 1
}
