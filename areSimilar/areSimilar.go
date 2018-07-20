package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(areSimilar([]int{
		832, 570, 570, 998, 533, 561, 455, 147, 894, 279,
		832, 998, 148, 570, 533, 561, 455, 147, 894, 279,
		832, 998, 148, 570, 533, 561, 894, 147, 455, 279,
		832, 998, 148, 570, 533, 561, 894, 147, 455, 279,
		832, 998, 148, 570, 533, 561, 894, 147, 455, 279,
		832, 998, 148, 148, 533, 561, 894, 147, 455, 279,
	}, []int{
		832, 570, 148, 998, 533, 561, 455, 147, 894, 279,
		832, 998, 148, 570, 533, 561, 455, 147, 894, 279,
		832, 998, 148, 570, 533, 561, 894, 147, 455, 279,
		832, 998, 148, 570, 533, 561, 894, 147, 455, 279,
		832, 998, 148, 570, 533, 561, 894, 147, 455, 279,
		832, 998, 148, 570, 533, 561, 894, 147, 455, 279,
	}))
}

func areSimilar(a []int, b []int) bool {
	discrepancies := 0
	inArray := func(needle int, haystack []int, base []int) (bool, int) {
		for i, val := range haystack {
			if needle == val && haystack[i] != base[i] {
				fmt.Println("Array Search Results:", needle, val, i)
				return true, i
			}
		}

		return false, -1
	}

	if equal := reflect.DeepEqual(a, b); !equal {
		for i := range a {
			if discrepancies > 1 {
				break
			}

			if a[i] != b[i] {
				discrepancies++
				is, idx := inArray(a[i], b, a)
				fmt.Println(is, idx)
				if !is {
					return false
				}

				b[i] = b[i] + b[idx]
				b[idx] = b[i] - b[idx]
				b[i] = b[i] - b[idx]
			}
		}

		return reflect.DeepEqual(a, b) && discrepancies <= 1
	}

	return true
}
