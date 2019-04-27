package algorithms

func twoSum(nums []int, target int) []int {
	for i, o_v := range nums {
		for k, in_v := range nums {
			if i == k {
				continue
			}
			if o_v + in_v == target {
				rs := []int{i, k}
				return rs
			}
		}
	}
	return []int{}
}