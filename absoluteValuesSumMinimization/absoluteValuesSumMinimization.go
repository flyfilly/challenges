package main

import (
	"fmt"
)

func main() {
	fmt.Println(absoluteValuesSumMinimization([]int{
		2, 4, 7,
	}))
}

func absoluteValuesSumMinimization(a []int) int {
	middleIdx := len(a) / 2

	if len(a)%2 == 0 {
		return a[middleIdx-1]
	} else {
		return a[middleIdx]
	}
}
