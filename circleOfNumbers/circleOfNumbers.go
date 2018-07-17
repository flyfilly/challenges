package main

import "fmt"

func circleOfNumbers(n int, firstNumber int) int {
	mid := n / 2

	if firstNumber >= mid {
		return firstNumber - mid
	} else {
		return mid + firstNumber
	}
}

func main() {
	fmt.Println(circleOfNumbers(10, 2))
}
