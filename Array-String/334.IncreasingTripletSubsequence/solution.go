package array_string

import "math"

func increasingTriplet(nums []int) bool {
    
	if len(nums)<3 {
		return false
	}
	firtMax:=nums[0]
	secoundMax:=math.MaxInt

	for i := 1; i < len(nums); i++ {
		if nums[i] > secoundMax {
			return true
		}else if nums[i] > firtMax {
			secoundMax = nums[i]
		}else{
			firtMax = nums[i]
		}
	}
	
	return false
}