package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums), len(nums))

	multiplier := 1
	for i := 0; i < len(nums); i++ {
		res[i] = multiplier
		multiplier *= nums[i]
		// println(res[i], multiplier)
	}

	multiplier = 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= multiplier
		multiplier *= nums[i]
		// println(res[i], multiplier)
	}

	return res
}
