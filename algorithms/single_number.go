package algorithms

// map[int]int
func singleNumber(nums []int) int {
	c := map[int]int{}

	for _, num := range nums {
		_, ok := c[num]
		if ok {
			c[num]++
		} else {
			c[num] = 1
		}
	}

	for i, n := range c {
		if n == 1 {
			return i
		}
	}
	return 0
}