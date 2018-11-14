package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(chessKnight("c2"))
}

func chessKnight(cell string) int {
	rows := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8}
	coords := strings.Split(cell, "")
	row := coords[0]
	col, _ := strconv.Atoi(coords[1])
	validMoves := 0
	getValidMoves := func(spaces int) int {
		moves := 0
		if (col - spaces) > 0 {
			moves++
		}
		if (col + spaces) < 9 {
			moves++
		}
		return moves
	}

	if (rows[row] - 2) > 0 {
		validMoves += getValidMoves(1)
	}

	if (rows[row] - 1) > 0 {
		validMoves += getValidMoves(2)
	}

	if (rows[row] + 2) < 9 {
		validMoves += getValidMoves(1)
	}

	if (rows[row] - 1) < 9 {
		validMoves += getValidMoves(2)
	}

	return validMoves
}
