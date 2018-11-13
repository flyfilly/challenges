package main

import (
	"fmt"
)

func main() {
	image := rotateImage([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	fmt.Println(image)
}

func rotateImage(a [][]int) [][]int {
	temp := 0
	n := len(a)
	for i := 0; i < n; i++ {
		for j := i; j < len(a[0]); j++ {
			temp = a[i][j]
			a[i][j] = a[j][i]
			a[j][i] = temp
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			temp = a[i][j]
			a[i][j] = a[i][n-1-j]
			a[i][n-1-j] = temp
		}
	}

	return a
}
