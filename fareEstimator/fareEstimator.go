package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(fareEstimator(30, 7, []float32{0.2, 0.35, 0.4, 0.45}, []float32{1.1, 1.8, 2.3, 3.5}))
}

func fareEstimator(ride_time int, ride_distance int, cost_per_minute []float32, cost_per_mile []float32) []float32 {
	fareEstimate := make([]float32, 0)

	for i := 0; i < len(cost_per_mile); i++ {
		estimate := math.Round(float64((cost_per_minute[i]*float32(ride_time))+(cost_per_mile[i]*float32(ride_distance)))*100) / 100
		fareEstimate = append(fareEstimate, float32(estimate))
	}

	return fareEstimate
}
