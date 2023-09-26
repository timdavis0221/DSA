package dsa

import (
	"regexp"
	"strings"
)

// ValidPalindrome recognizes whether a given string is a valid palindrome.
func ValidPalindrome(s string) bool {
	// Remove all non-alphanumeric characters and convert all uppercase letters into lowercase letters.
	pattern := regexp.MustCompile(`[^0-9A-Za-z]+`)
	alphanumericStr := pattern.ReplaceAllString(s, "")
	convertedStr := strings.ToLower(alphanumericStr)
	// Use two pointers to compare start character with end character from start index and end index.
	end := len(convertedStr) - 1
	for start := range convertedStr {
		if convertedStr[start] != convertedStr[end] {
			return false
		}
		end--
	}
	return true
}
