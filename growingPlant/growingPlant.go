package main

import (
	"fmt"
)

func main() {
	fmt.Println(growingPlant(5, 2, 7))
}

func growingPlant(upSpeed int, downSpeed int, desiredHeight int) int {
	days := 0
	height := 0

	for {
		days++
		height += upSpeed
		if height < desiredHeight {
			height -= downSpeed
		} else {
			break
		}
	}

	return days
}
