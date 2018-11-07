package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(knapsackLight(15, 2, 20, 3, 2)) //15
}

func knapsackLight(value1 int, weight1 int, value2 int, weight2 int, maxW int) int {
	switch {
	case (weight1 + weight2) <= maxW:
		return value1 + value2
	case weight1 <= maxW && weight2 <= maxW:
		return int(math.Max(float64(value1), float64(value2)))
	case weight1 <= maxW:
		return value1
	case weight2 <= maxW:
		return value2
	}

	return 0
}
