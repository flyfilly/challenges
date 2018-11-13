package main

import (
	"fmt"
)

func main() {
	fmt.Println(electionsWinners([]int{5, 1, 3, 4, 1}, 0))
}

func electionsWinners(votes []int, k int) int {
	high := 0
	highExists := 0
	couldWin := 0

	for _, vote := range votes {
		if high < vote {
			high = vote
		}
	}

	for _, vote := range votes {
		switch {
		case high < (vote + k):
			couldWin++
		case high == vote:
			highExists++
		}
	}

	if 0 == k && highExists == 1 {
		return highExists
	}

	return couldWin
}
