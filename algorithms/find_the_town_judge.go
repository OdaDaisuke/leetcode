package algorithms

import "strconv"

func findJudge(N int, trust [][]int) int {
	if len(trust) == 0 {
		return N
	}
	dic := map[string]int{}
	for _, n := range trust {
		k := strconv.Itoa(n[1])
		dic[k] = dic[k] + 1
	}

	var maxKey string
	var maxVal int
	for k, v := range dic {
		if maxVal < v {
			maxVal = v
			maxKey = k
		}
	}

	if maxVal >= N - 1 {
		r, _ := strconv.Atoi(maxKey)
		for _, f := range trust {
			if f[0] == r {
				return -1
			}
		}
		return r
	}
	return -1
}