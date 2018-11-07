package main

import (
	"fmt"
)

func main() {
	fmt.Println(extractEachKth([]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}, 3))

	fmt.Println(extractEachKth([]int{
		1, 1, 1, 1, 1,
	}, 1))

	fmt.Println(extractEachKth([]int{
		1, 2, 1, 2, 1, 2, 1, 2,
	}, 2))

	fmt.Println(extractEachKth([]int{
		1, 2, 1, 2, 1, 2, 1, 2,
	}, 10))
}

func extractEachKth(inputArray []int, k int) []int {
	arr := make([]int, 0)
	for i := 0; i < len(inputArray); i++ {
		if (i+1)%k != 0 {
			arr = append(arr, inputArray[i])
		}
	}

	return arr
}
