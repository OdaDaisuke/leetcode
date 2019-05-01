package algorithms

func reverse(x int) int {
	var isMinus bool
	if x < 0 {
		isMinus = true
		x = -x
	}

	sx := []byte(strconv.Itoa(x))
	sxLen := len(sx)
	for i := 0; i < sxLen / 2; i++ {
		sx[i], sx[sxLen - i - 1] = sx[sxLen - i - 1], sx[i]
	}
	intX, _ := strconv.Atoi(string(sx))

	if -intX < -2147483648 || intX >= 2147483648 {
		return 0
	}

	if isMinus {
		return -intX
	}

	return intX
}
