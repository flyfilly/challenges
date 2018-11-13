package main

import (
	"fmt"
)

func main() {
	fmt.Println(isBeautifulString("bbbaacdafe"))
}

func isBeautifulString(inputString string) bool {
	var makeup [26]int

	for _, chr := range inputString {
		makeup[chr-'a']++
	}

	for i := 1; i < len(makeup); i++ {
		if makeup[i] > makeup[i-1] {
			return false
		}
	}

	return true
}

// func isBeautifulString(inputString string) bool {
// 	makeup := make(map[rune]int)

// 	for _, chr := range inputString {
// 		makeup[chr]++
// 	}

// 	keys := reflect.ValueOf(makeup).MapKeys()
// 	sort.Slice(keys, func(i, j int) bool {
// 		return keys[i].Interface().(int32) < keys[j].Interface().(int32)
// 	})

// 	for i := 0; i < (len(keys) - 1); i++ {
// 		here := rune(keys[i].Interface().(int32))
// 		next := rune(keys[i+1].Interface().(int32))

// 		fmt.Println(makeup[here], makeup[next])

// 		if makeup[here] < makeup[next] {
// 			return false
// 		}
// 	}

// 	return true
// }
