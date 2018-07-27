package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(avoidObstacles([]int{2, 3}))
}

func avoidObstacles(inputArray []int) int {
	length := 1
	inArray := func(needle int, haystack []int) bool {
		fmt.Println("Searching for", needle, "in", haystack)
		for _, num := range haystack {
			if needle == num {
				fmt.Println("found", needle, "in", haystack)
				return true
			}
		}
		return false
	}

	sort.Ints(inputArray)
	inc := length

	for i := 0; i < inputArray[len(inputArray)-1]+1; i++ {
		if inArray(inc, inputArray) {
			length++
			i = 0
			inc = length
		} else {
			inc += length
		}
	}

	return length
}
