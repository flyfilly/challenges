package main

import (
	"fmt"
)

func main() {
	fmt.Println(matrixElementsSum([][]int{
		{1, 1, 1, 0},
		{0, 5, 0, 1},
		{2, 1, 3, 10},
	}))
}

func matrixElementsSum(matrix [][]int) int {
	sum := 0
	var unsutables []int
	isSuitable := func(val int) bool {
		good := true

		for _, item := range unsutables {
			if item == val {
				good = false
				break
			}
		}

		return good
	}

	for i, arr := range matrix {
		for j, _ := range arr {
			if matrix[i][j] == 0 {
				unsutables = append(unsutables, j)
			}

			if i > 0 {
				if matrix[i-1][j] > 0 {
					if isSuitable(j) {
						sum += matrix[i][j]
					}
				}
			} else {
				sum += matrix[i][j]
			}
		}
	}

	return sum
}
