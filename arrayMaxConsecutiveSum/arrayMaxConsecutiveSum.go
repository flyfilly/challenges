package main

import "fmt"

func main() {
	fmt.Println(arrayMaxConsecutiveSum([]int{
		2, 3, 5, 1, 6,
	}, 2))
}

func arrayMaxConsecutiveSum(inputArray []int, k int) int {
	max := 0

	for i := 0; i < len(inputArray); i++ {
		until := (i + k - 1)
		sum := 0

		if until < len(inputArray) {
			for j := i; j <= until; j++ {
				sum += inputArray[j]
			}

			if sum > max {
				max = sum
			}
		}
	}

	return max
}
