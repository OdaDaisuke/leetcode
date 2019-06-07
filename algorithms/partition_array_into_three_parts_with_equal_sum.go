package algorithms

func canThreePartsEqualSum(A []int) bool {
	var answer int
	var tmpSum int
	var addCount int

	for _, a := range A {
		answer = answer + a
	}
	answer = answer / 3

	for _, a := range A {
		tmpSum = tmpSum + a
		if tmpSum == answer {
			tmpSum = 0
			addCount++
		}
		if addCount == 3 {
			return true
		}
	}

	return false
}