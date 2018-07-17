package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isLucky(1230))
}

func isLucky(n int) bool {
	str := strconv.Itoa(n)
	half := len(str) / 2
	first, second := str[0:half], str[half:]
	sumFirst, sumSecond := 0, 0

	for i := 0; i < half; i++ {
		val1, _ := strconv.Atoi(string(first[i]))
		val2, _ := strconv.Atoi(string(second[i]))

		sumFirst += val1
		sumSecond += val2
	}

	return sumFirst == sumSecond
}
