package longest_consecutive_sequence

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	count := 1
	check := 0

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			continue
		}
		if nums[i]+1 == nums[i+1] {
			count++
			// println(count)
		} else {
			// fmt.Println("check ", check, " count", count)
			if count >= check {
				check = count
			}
			count = 1
		}
	}
	if check > count {
		return check
	}
	// fmt.Println(nums)
	return count
}
