package main

import "fmt"

func main() {
	fmt.Println(arrayReplace([]int{
		1, 2, 1,
	}, 1, 3))
}

func arrayReplace(inputArray []int, elemToReplace int, substitutionElem int) []int {
	newArr := []int{}
	for _, val := range inputArray {
		if val == elemToReplace {
			newArr = append(newArr, substitutionElem)
		} else {
			newArr = append(newArr, val)
		}
	}

	return newArr
}
