package group_anagram

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)

	for _, v := range strs {
		// rv := []rune(v)
		// sort.Slice(rv, func(i, j int) bool {
		// 	return v[i] < v[j]
		// })
		sArr := strings.Split(v, "")
		sort.Strings(sArr)
		s := strings.Join(sArr, "")
		// println(val)
		// println(v)
		strMap[s] = append(strMap[s], v)
	}
	// fmt.Printf("intMap: %v\n", strMap)
	var arr [][]string

	for _, v := range strMap {
		arr = append(arr, v)
	}

	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})

	return arr
}
