package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isDigit("@"))
}

func isDigit(symbol string) bool {
	_, error := strconv.Atoi(symbol)
	return error == nil
}
