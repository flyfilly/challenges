package main

import (
	"fmt"
)

func main() {
	fmt.Println(buildPalindrome("abcdc"))
}

func buildPalindrome(st string) string {
	var canConvert bool
	for i := len(st); ; i++ {
		canConvert = true
		for j := 0; j < i-j-1; j++ {
			if i-j <= len(st) && st[j] != st[i-j-1] {
				canConvert = false
				break
			}
		}
		if canConvert {
			for j := len(st); j < i; j++ {
				st += string(st[i-j-1])
			}
			return st
		}
	}
}
