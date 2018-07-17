package main

import (
	"fmt"
)

func main() {
	fmt.Println(sortByHeight([]int{
		-1, 150, 190, 170, -1, -1, 160, 180,
	}))
}

func sortByHeight(a []int) []int {
	tmp := make([]int, 0)
	sort := func(ints []int) []int {
		for i := 0; i < len(ints); i++ {
			for j := 0; j < len(ints); j++ {
				if ints[i] < ints[j] {
					ints[i] = ints[i] + ints[j]
					ints[j] = ints[i] - ints[j]
					ints[i] = ints[i] - ints[j]
				}
			}
		}

		return ints
	}

	for _, num := range a {
		if num != -1 {
			tmp = append(tmp, num)
		}
	}

	sort(tmp)

	for i, num := range a {
		if num != -1 {
			x := tmp[0]
			tmp = tmp[1:]

			a[i] = x
		}
	}

	return a
}
