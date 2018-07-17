package main

import (
	"fmt"
)

func candies(n int, m int) int {
	return (m / n) * n
}

func main() {
	fmt.Println(candies(3, 10))
}
