package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evenDigitsOnly(642386))
}

func evenDigitsOnly(n int) bool {
	num := strconv.Itoa(n)
	for i := 0; i < len(num); i++ {
		if digit, err := strconv.Atoi(string(num[i])); err == nil {
			if digit%2 != 0 {
				return false
			}
		}
	}
	return true
}
