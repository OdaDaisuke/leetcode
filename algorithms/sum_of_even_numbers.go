package algorithms

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	var rs = make([]int, len(A))
	for qi, q := range queries {
		tmpA := A
		tarIdx := q[1]
		val := q[0]

		tmpA[tarIdx] = tmpA[tarIdx] + val
		for _, v := range tmpA {
			if v % 2 == 0 {
				rs[qi] = rs[qi] + v
			}
		}
	}

	return rs
}