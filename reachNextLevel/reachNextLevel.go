package main

import "fmt"

func reachNextLevel(experience int, threshold int, reward int) bool {
	return (experience + reward) >= threshold
}

func main() {
	fmt.Println(reachNextLevel(80, 199, 15))
}
