package main

import (
	"fmt"
)

func main() {
	fmt.Println(boxBlur([][]int{
		{1, 1, 1},
		{1, 7, 1},
		{1, 1, 1},
	}))
}

func boxBlur(image [][]int) [][]int {
	blurred := make([][]int, 0)

	for y := 0; y < len(image)-2; y++ {
		row := make([]int, 0)

		for x := 0; x < len(image[y])-2; x++ {
			sum := 0

			for i := y; i < y+3; i++ {
				for j := x; j < x+3; j++ {
					sum += image[i][j]
				}
			}
			row = append(row, sum/9)
		}
		blurred = append(blurred, row)
	}
	return blurred
}
