package main

func main() {
	arrayMaximalAdjacentDifference([]int{2, 4, 1, 0})
}

func arrayMaximalAdjacentDifference(inputArray []int) int {
	maxDiff := 0

	for i := 1; i < len(inputArray); i++ {
		prev, curr := inputArray[i-1], inputArray[i]
		diff := prev - curr

		if diff < 0 {
			diff = -diff
		}

		if diff > maxDiff {
			maxDiff = diff
		}
	}

	return maxDiff
}
