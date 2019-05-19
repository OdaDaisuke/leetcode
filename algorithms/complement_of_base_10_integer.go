package algorithms

import (
	"fmt"
	"strconv"
)

func bitwiseComplement(N int) int {
	var twoD string
	b := fmt.Sprintf("%b", N)
	sb := string(b)
	sl := len(sb)
	for i := 0; i < sl; i++ {
		if string(sb[i]) == "0" {
			twoD = twoD + "1"
		} else {
			twoD = twoD + "0"
		}
	}
	rs, _ := strconv.ParseInt(twoD, 2, 0)
	return int(rs)
}