package main

import (
	"fmt"
	"strconv"
)

/* Example

For n = 240, the output should be lateRide(n) = 4.

Since 240 minutes have passed, the current time is 04:00.
The digits sum up to 0 + 4 + 0 + 0 = 4, which is the answer.

For n = 808, the output should be
lateRide(n) = 14.

808 minutes mean that it's 13:28 now, so the answer should be
1 + 3 + 2 + 8 = 14.*/

func lateRide(n int) int {
	hrs := n / 60
	mins := n % 60
	time := []byte(strconv.Itoa(hrs) + strconv.Itoa(mins))
	out := 0

	for _, v := range time {
		val, _ := strconv.Atoi(string(v))
		out += val
	}

	return out
}

func main() {
	fmt.Println(lateRide(808))
}
