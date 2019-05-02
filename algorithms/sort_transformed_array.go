package algorithms

func sortTransformedArray(nums []int, a int, b int, c int) []int {
	var rs []int
	for _, num := range nums {
		calc := a * num * num + num * b + c
		rs = append(rs, calc)
	}

	l := len(rs) - 1
	for i := 0; i < l; i++ {
		for j, n := range rs {
			if j > 0 && rs[j - 1] > n {
				rs[j], rs[j - 1] = rs[j - 1], n
			}
		}
	}

	return rs
}