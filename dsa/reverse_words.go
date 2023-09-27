package dsa

import (
	"strings"
)

// ReverseWords reverses the order of the words in the given string.
// Note:
// A word is defined as as a sequence of non-space characters, which will be separated by at least one space.
// Input string s may contain leading or trailing spaces or multiple spaces between words.
// The returned string should only have a single space between the words.
func ReverseWords(s string) string {
	// First thought:
	// Split the string then use the two pointer to swap characters.
	newStr := strings.Fields(s)
	start, end := 0, len(newStr)-1
	for start < end {
		newStr[start], newStr[end] = newStr[end], newStr[start]
		start++
		end--
	}

	var resultStr string
	for i, v := range newStr {
		if i < len(newStr)-1 {
			resultStr = resultStr + v + " "
		} else {
			resultStr += v
		}
	}
	return resultStr
}
