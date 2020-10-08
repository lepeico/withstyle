package withstyle

import (
	"math"
)

func intRound(x float64) int {
	return int(math.Round(float64(x)))
}

func findMinAndMax(xs ...float64) (min float64, max float64) {
	min = xs[0]
	max = xs[0]
	for _, value := range xs {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func min(xs ...float64) (min float64) {
	min, _ = findMinAndMax(xs...)
	return
}

func max(xs ...float64) (max float64) {
	_, max = findMinAndMax(xs...)
	return
}
