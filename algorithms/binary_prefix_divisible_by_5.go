package algorithms

func prefixesDivBy5(A []int) []bool {
	num, answer := 0, make([]bool, len(A))
	for i, a := range A {
		num = (num * 2 + a) % 5

		if num == 0 {
			answer[i] = true
		}
	}
	return answer
}