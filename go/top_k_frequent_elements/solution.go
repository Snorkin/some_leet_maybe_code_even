package top_k_frequent_elements

import "sort"

func topKFrequent(nums []int, k int) []int {
	numsMap := make(map[int]int)

	for _, v := range nums {
		// println(v)
		numsMap[v]++
	}

	res := make([]int, 0, len(numsMap))

	for key := range numsMap {
		res = append(res, key)
	}

	sort.SliceStable(res, func(i, j int) bool {
		return numsMap[res[i]] > numsMap[res[j]]
	})

	// fmt.Printf("Map: %v\n", numsMap)
	return res[:k]
}
