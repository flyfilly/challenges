package main

import (
	"fmt"
)

/*
Some phone usage rate may be described as follows:

first minute of a call costs min1 cents,
each minute from the 2nd up to 10th (inclusive) costs min2_10 cents
each minute after 10th costs min11 cents.
You have s cents on your account before the call. What is the
duration of the longest call (in minutes rounded down to the
nearest integer) you can have?

EXAMPLE:
For min1 = 3, min2_10 = 1, min11 = 2, and s = 20,
the output should be
phoneCall(min1, min2_10, min11, s) = 14. */
func phoneCall(min1 int, min2_10 int, min11 int, s int) int {
	cost := min1
	minutes := 0

	for ok := (cost <= s); ok; ok = (cost <= s) {
		s -= cost
		minutes++
		if minutes == 1 {
			cost = min2_10
		} else if minutes == 10 {
			cost = min11
		}
	}

	return minutes
}

func main() {
	fmt.Println(phoneCall(3, 1, 2, 20))
}
