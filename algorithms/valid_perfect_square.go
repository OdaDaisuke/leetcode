package algorithms

import "math"

func isPerfectSquare(num int) bool {
	s := int(math.Sqrt(float64(num)))
	if s * s == num {
		return true
	}
	return false
}