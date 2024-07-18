package array_string

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	totalProduct := 1
	zeroIndex := -1
	answer := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if zeroIndex != -1 {
				return answer
			}
			zeroIndex = i
		} else {
			totalProduct *= nums[i]
		}
	}

	if zeroIndex != -1 {
		answer[zeroIndex] = totalProduct
		return answer
	}

	for i := 0; i < len(nums); i++ {
		answer[i] = totalProduct / nums[i]
	}
	return answer
}