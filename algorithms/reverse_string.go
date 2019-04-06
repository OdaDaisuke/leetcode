package algorithms

func reverseString(s []byte)  {
	sLen := len(s) - 1
	loopLen := len(s) / 2
	for i := 0; i < loopLen; i++ {
		s[i], s[sLen - i] = s[sLen - i], s[i]
	}
}

// https://leetcode.com/submissions/detail/220435049/