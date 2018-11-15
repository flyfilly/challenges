package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(sumUpNumbers("2 apples, 12 oranges"))
}

func sumUpNumbers(inputString string) int {
	regex, _ := regexp.Compile("[^0-9 ]+")
	clean := regex.ReplaceAllString(inputString, " ")
	parts := strings.Split(clean, " ")
	sum := 0

	fmt.Println(parts)

	for _, part := range parts {
		if num, err := strconv.Atoi(part); err == nil {
			sum += num
		}
	}

	return sum
}
