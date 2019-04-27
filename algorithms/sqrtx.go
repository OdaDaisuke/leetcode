package algorithms

import "math"

func mySqrt(x int) int {
	c := math.Sqrt(float64(x))
	return int(c)
}