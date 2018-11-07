package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println(firstDigit("var_1_Int"))
	fmt.Println(firstDigit("_Aas_23"))
}

func firstDigit(inputString string) string {
	returnStr := ""
	for _, char := range inputString {
		if digit, err := strconv.Atoi(string(char)); err == nil {
			returnStr = strconv.Itoa(digit)
			break
		}
	}
	return returnStr
}
