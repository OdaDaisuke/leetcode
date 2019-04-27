package algorithms

import (
	"fmt"
	"strings"
	"strconv"
)

func complexNumberMultiply(a string, b string) string {
	as := strings.Split(a, "+")
	bs := strings.Split(b, "+")

	aF, _ := strconv.Atoi(as[0])
	aL, _ := strconv.Atoi(as[1][:len(as[1])-1])
	bF, _ := strconv.Atoi(bs[0])
	bL, _ := strconv.Atoi(bs[1][:len(bs[1])-1])

	rf := aF * bF + -(aL * bL)
	rl := aF * bL + aL * bF

	return fmt.Sprintf("%d+%di", rf, rl)
}
