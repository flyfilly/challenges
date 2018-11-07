package main

import (
	"fmt"
)

func main() {
	fmt.Println(differentSymbolsNaive("aba"))
}

func differentSymbolsNaive(s string) int {
	m := make(map[rune]int)

	for _, char := range s {
		m[char]++
	}

	return len(m)
}
