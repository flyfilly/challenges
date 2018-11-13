package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findEmailDomain("prettyandsimple@example.com"))
}

func findEmailDomain(address string) string {
	parts := strings.Split(address, "@")
	return parts[len(parts)-1]
}
