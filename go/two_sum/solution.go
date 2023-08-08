package two_sum

func twoSum(nums []int, target int) []int {
	intMap := make(map[int]int)
	for id, value := range nums {
		// println(id)

		diff := target - value
		// println(diff)
		if v, check := intMap[diff]; check {
			// println("check " + fmt.Sprintln(v))
			return []int{v, id}
		}
		intMap[value] = id
	}
	// fmt.Printf("intMap: %v\n", intMap)
	return []int{0}
}
