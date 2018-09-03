package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(chessBoardCellColor("A1", "H3"))
}

func chessBoardCellColor(cell1 string, cell2 string) bool {
	board := "ABCDEFGH"
	colA := strings.Index(board, string(cell1[0]))%2 == 0
	a, _ := strconv.Atoi(string(cell1[1]))
	rowA := a%2 == 0

	colB := strings.Index(board, string(cell2[0]))%2 == 0
	b, _ := strconv.Atoi(string(cell2[1]))
	rowB := b%2 == 0

	if colA == colB {
		return rowA == rowB
	}
	return rowA != rowB
}
