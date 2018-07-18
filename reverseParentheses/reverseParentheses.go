package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseParentheses("a(bcdefghijkl(mno)p)q"))
	fmt.Println(reverseParentheses("abc(cba)ab(bac)c"))
}

func reverseParentheses(s string) string {
	brackets := make([][]int, 0)

	reverse := func(str string) string {
		revstr := ""

		for i := len(str) - 1; i >= 0; i-- {
			revstr += string(str[i])
		}

		return revstr
	}

	for i, char := range s {
		switch {
		case string(char) == "(":
			brackets = append(brackets, []int{i})
		case string(char) == ")":
			for j := len(brackets) - 1; j >= 0; j-- {
				if len(brackets[j]) < 2 {
					brackets[j] = append(brackets[j], i)
					break
				}
			}
		}
	}

	for i := len(brackets) - 1; i >= 0; i-- {
		start, end := brackets[i][0], brackets[i][1]+1
		s = s[:start] + reverse(s[start:end]) + s[end:]
	}

	return strings.Replace(strings.Replace(s, ")", "", -1), "(", "", -1)
}
