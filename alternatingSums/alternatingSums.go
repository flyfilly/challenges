package main

import "fmt"

func main() {
	fmt.Println(alternatingSums([]int{
		50, 60, 60, 45, 70,
	}))
}

func alternatingSums(a []int) []int {
	team1, team2 := 0, 0
	for i, mbr := range a {
		if i%2 == 0 {
			team1 += mbr
		} else {
			team2 += mbr
		}
	}

	return []int{team1, team2}
}
