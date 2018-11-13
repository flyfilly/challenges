package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(lineEncoding("abbcabb"))
}

func lineEncoding(s string) string {
	encoded := ""
	len := len(s)
	recurrences := 1

	for i := 0; i < len; i++ {
		if !((i + 1) == len) {
			if s[i] == s[(i+1)] {
				recurrences++
			} else {
				if recurrences > 1 {
					encoded += strconv.Itoa(recurrences) + string(s[i])
				} else {
					encoded += string(s[i])
				}
				recurrences = 1
			}
		} else {
			if recurrences == 1 {
				encoded += string(s[i])
			} else {
				encoded += strconv.Itoa(recurrences) + string(s[i])
			}
		}
	}

	return encoded
}
