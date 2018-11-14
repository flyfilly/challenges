package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(deleteDigit(152))
}

func deleteDigit(n int) int {
	high := 0
	str := strconv.Itoa(n)

	for i, _ := range str {
		num, _ := strconv.Atoi(str[:i] + str[i+1:])
		if num > high {
			high = num
		}
	}
	return high
}
