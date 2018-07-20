package main

import "fmt"

func main() {
	fmt.Println(arrayChange([]int{2, 1, 10, 1}))
}

func arrayChange(inputArray []int) int {
	changes := 0

	for i := 1; i < len(inputArray); i++ {
		prev, curr := inputArray[i-1], inputArray[i]

		if prev >= curr {
			needsToBe := prev + 1
			changes += needsToBe - curr
			inputArray[i] = needsToBe
		}
	}

	return changes
}
