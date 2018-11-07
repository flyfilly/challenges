package main

import (
	"fmt"
)

func main() {
	fmt.Println(stringsRearrangement([]string{
		"aba",
		"bbb",
		"bab",
	}))
}

func check(start int, inputArray []string) bool {
	swap := func(i int, j int) {
		tmp := inputArray[i]
		inputArray[i] = inputArray[j]
		inputArray[j] = tmp
	}
	if start == len(inputArray) {
		for i := 0; i < len(inputArray)-1; i++ {
			var differences = 0
			for j := 0; j < len(inputArray[i]); j++ {
				if inputArray[i][j] != inputArray[i+1][j] {
					differences++
				}
			}
			if differences != 1 {
				return false
			}
		}
		return true
	}
	for i := start; i < len(inputArray); i++ {
		swap(start, i)
		if check(start+1, inputArray) {
			return true
		}
		swap(start, i)
	}
	return false
}

func stringsRearrangement(inputArray []string) bool {
	return check(0, inputArray)
}
