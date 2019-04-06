package algorithms

func sortArrayByParity(A []int) []int {
	var even []int
	var odds []int
	for _, n := range A {
		if n % 2 == 0 {
			even = append(even, n)
		} else {
			odds = append(odds, n)
		}
	}

	return append(even, odds...)
}