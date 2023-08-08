package valid_anagram

func isAnagram(s string, t string) bool {
	// - my initial solution
	// rs := []rune(s)
	// rt := []rune(t)

	// sort.Slice(rs, func(i, j int) bool {

	// 	return rs[i] < rs[j]
	// })
	// sort.Slice(rt, func(i, j int) bool {

	// 	return rt[i] < rt[j]
	// })

	// return string(rs) == string(rt)

	// - gigabrain optimized solution
	chars := make(map[rune]int, 26) // eng letters count
	for _, v := range s {
		chars[v]++
	}

	for _, v := range t {
		chars[v]--
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}

	}
	return true
}
