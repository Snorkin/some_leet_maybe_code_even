package contains_duplicate

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})

	for _, v := range nums {
		if _, check := m[v]; check {
			return true
		}
		m[v] = struct{}{}
	}

	return false
}
