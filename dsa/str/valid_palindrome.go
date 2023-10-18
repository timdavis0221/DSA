package str

import (
	"strings"
	"unicode"
)

// ValidPalindrome recognizes whether a given string is a valid palindrome.
func ValidPalindrome(s string) bool {
	// Remove all non-alphanumeric characters and convert all uppercase letters into lowercase letters.
	runeSlice := make([]rune, 0)
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			runeSlice = append(runeSlice, v)
		}
	}
	convertedStr := strings.ToLower(string(runeSlice))
	// Use two pointers to compare start character with end character from start index and end index.
	start := 0
	end := len(convertedStr) - 1
	for start < end {
		if convertedStr[start] != convertedStr[end] {
			return false
		}
		start++
		end--
	}
	return true
}
