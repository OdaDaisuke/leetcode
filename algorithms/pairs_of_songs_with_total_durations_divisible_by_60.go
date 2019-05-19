package algorithms

func numPairsDivisibleBy60(time []int) int {
	var c int
	l := len(time)
	for i, t := range time {
		for j := i + 1; j < l; j++ {
			if (t + time[j]) % 60 == 0 {
				c++
			}
		}
	}

	return c
}