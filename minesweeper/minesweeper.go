package main

import "fmt"

func main() {
	fmt.Println(minesweeper([][]bool{
		{true, false, false},
		{false, true, false},
		{false, false, false},
	}))
}

func minesweeper(matrix [][]bool) [][]int {
	grid := make([][]int, 0)
	for y := 0; y < len(matrix); y++ {
		row := make([]int, 0)
		for x := 0; x < len(matrix[y]); x++ {
			bombs := 0

			if y-1 >= 0 && x-1 >= 0 && matrix[y-1][x-1] {
				bombs++
			}

			if x-1 >= 0 && matrix[y][x-1] {
				bombs++
			}

			if y+1 < len(matrix) && x-1 >= 0 && matrix[y+1][x-1] {
				bombs++
			}

			if y-1 >= 0 && matrix[y-1][x] {
				bombs++
			}

			if y+1 < len(matrix) && matrix[y+1][x] {
				bombs++
			}

			if y-1 >= 0 && x+1 < len(matrix[y]) && matrix[y-1][x+1] {
				bombs++
			}

			if x+1 < len(matrix[y]) && matrix[y][x+1] {
				bombs++
			}

			if y+1 < len(matrix) && x+1 < len(matrix[y]) && matrix[y+1][x+1] {
				bombs++
			}

			row = append(row, bombs)
		}
		grid = append(grid, row)
	}
	return grid
}
