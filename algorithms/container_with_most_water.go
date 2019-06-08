package algorithms

/*
1. row * colでloop
2. surface計算、最大値をget
3. 2をreturn
*/

func maxArea(height []int) int {
	var max int

	for hi, h := range height {
		for hi2, h2 := range height {
			if hi2 == hi || hi2 < hi {
				continue
			}

			// height diff * axios diff
			higher := h2
			lower := h
			if lower > higher {
				higher, lower = h, h2
			}
			absoluteHeight := higher - (higher - lower)
			axiosDiff := hi2 - hi
			if axiosDiff < 0 {
				axiosDiff = axiosDiff * -1
			}

			sum := axiosDiff * absoluteHeight
			if sum > max {
				max = sum
			}
		}
	}

	return max
}