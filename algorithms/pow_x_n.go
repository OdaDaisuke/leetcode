package algorithms

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	var rs float64
	rs = myPow(x, n / 2)
	if n % 2 == 0 {
		return rs * rs
	} else {
		if n > 0 {
			return x * rs * rs
		} else {
			return (rs * rs ) / x
		}
	}
}