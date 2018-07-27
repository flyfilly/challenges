package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(isIPv4Address("172.16.254.1"))
}

func isIPv4Address(inputString string) bool {
	parts := strings.Split(inputString, ".")

	if len(parts) < 4 {
		return false
	}

	for _, part := range parts {
		it, err := strconv.Atoi(part)

		if err != nil || it < 0 || it > 255 {
			return false
		}
	}

	return true
}
