package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(isMAC48Address("00-1B-63-84-45-E6"))
}

func isMAC48Address(inputString string) bool {
	match, _ := regexp.MatchString("^([0-9a-fA-F]{2}[-]){5}([0-9a-fA-F]{2})$", inputString)
	return match
}
