package main

import (
	"fmt"
	"strconv"
)

func largestNumber(n int) int {
	str := ""
	for i := 0; i < n; i++ {
		str += "9"
	}
	num, _ := strconv.ParseInt(str, 10, 64)

	return int(num)
}

func main() {
	fmt.Println(largestNumber(3))
}
