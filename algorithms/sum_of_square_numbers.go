package algorithms

import "math"

func judgeSquareSum(c int) bool {
	i := 0
	fc := int(math.Sqrt(float64(c)))
	for i <= fc {
		s := i * i + fc * fc
		if s == c {
			return true
		} else if s < c {
			i++
		} else {
			fc--
		}
	}
	return false
}