package algorithms

import "fmt"

func reverseStr(s string, k int) string {
	sLen := len(s)
	var r string

	for i := 0; i < sLen; i += k {
		head := (i / k) * k
		tail := head + k
		if tail > sLen {
			tail = sLen
		}

		curS := s[head:tail]
		curLen := len(curS)
		if i / k % 2 == 1 {
			r += curS
			continue
		}
		for _, v := range curS {
			curS = fmt.Sprintf("%c%s", v, curS)
		}
		curS = curS[:curLen]
		r += curS
	}

	return r
}