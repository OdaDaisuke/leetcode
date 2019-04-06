package algorithms

func numJewelsInStones(J string, S string) int {
	var count int

	for _, s := range S {
		for _, j := range J {
			if j == s {
				count++
			}
		}
	}

	return count
}
