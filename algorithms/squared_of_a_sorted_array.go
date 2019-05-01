package algorithms

func sortedSquares(A []int) []int {
	var squared []int
	for _, a := range A {
		squared = append(squared, a * a)
	}

	l := len(A)
	for k := 0; k < l - 1; k++ {
		for i, n := range squared {
			if i + 1 < l && n > squared[i + 1] {
				squared[i + 1], squared[i] = n, squared[i + 1]
			}
		}
	}

	return squared
}