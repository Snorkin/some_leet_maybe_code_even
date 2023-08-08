package valid_palindrome

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i := 0
	j := len(s) - 1

	for i < j {
		if !(unicode.IsDigit(rune(s[i])) || unicode.IsLetter(rune(s[i]))) {
			i++
			continue
		}

		if !(unicode.IsDigit(rune(s[j])) || unicode.IsLetter(rune(s[j]))) {
			j--
			continue
		}

		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}
