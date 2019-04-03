package algorithms

import "strconv"

func isPalindrome(x int) bool {
	loopLen := len(strconv.Itoa(x))/2 + 1
	strX := strconv.Itoa(x)
	xLen := len(strX)

	for i := 1; i < loopLen; i++ {
		v := strX[i-1]
		if v != strX[xLen-i] {
			return false
		}
	}

	return true
}
