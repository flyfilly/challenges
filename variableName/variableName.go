package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(variableName("_var_1__Int"))
}

func variableName(name string) bool {
	tested, _ := regexp.MatchString("^[a-zA-Z_][0-9a-zA-Z_]*$", name)
	return tested
}
