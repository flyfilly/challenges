package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(bishopAndPawn("a1", "c3"))
}

func bishopAndPawn(bishop string, pawn string) bool {
	x := math.Abs(float64(int(bishop[1]) - int(pawn[1])))
	y := math.Abs((float64(bishop[0]) - 96) - (float64(pawn[0]) - 96))

	return x == y
}
