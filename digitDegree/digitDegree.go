package main

import (
	"math"
)

func main() {

}

func digitDegree(n int) int {
	return add(n, 0)
}

func add(n int, m int) int {
	if n >= 10 {
		t := 0

		for n > 0 {
			t += n % 10
			n = int(math.Floor(float64(n) / float64(10)))
		}

		return add(t, (m + 1))
	}

	return m
}
