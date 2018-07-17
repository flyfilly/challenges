package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseParentheses("a(bc)de"))
}

func reverseParentheses(s string) string {
	start := strings.Index(s, "(")
	end := strings.Index(s, ")")
	reverse := func(str string) string {
		revstr := ""

		for i := len(str) - 1; i >= 0; i-- {
			revstr += string(str[i])
		}

		return revstr
	}

	return s[:start] + reverse(s[start+1:end]) + s[end+1:]
}
