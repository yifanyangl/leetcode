package algorithms

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		vTarget := target - v
		for iNext, vNext := range nums[i+1:] {
			if vTarget == vNext {
				return []int{i, iNext + i + 1}
			}
		}
	}
	return nil
}
