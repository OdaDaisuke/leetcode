package algorithms

func moveZeroes(nums []int)  {
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		if nums[i] == 0 {
			nums = calc(nums, i)
			continue
		}
	}
}

func calc(nums []int, i int) []int {
	nums = append(nums[:i], nums[i + 1:]...)
	nums = append(nums, 0)
	numsLen := len(nums) / 2
	var loopCount int

	for {
		if loopCount >= numsLen {
			break
		} else if nums[i] == 0 {
			nums = append(nums[:i], nums[i + 1:]...)
			nums = append(nums, 0)
			loopCount++
			continue
		}

		break
	}

	return nums
}

// https://leetcode.com/submissions/detail/220441545/