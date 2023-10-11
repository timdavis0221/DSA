package str

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

// ReverseWords2 uses in-place algorithm to reverse words for consuming less space purpose.
// O(1) if string is mutable, otherwise O(n).
func ReverseWords2(s string) string {
	// For modifying immutable string in-place.
	byteStr := []byte(s)
	// Slow pointer that used to store the filtered characters where its length is less than
	// the length of string (fast pointer).
	slow := 0
	// Fast pointer that used to track (remove, skip) the blank and convert the input string before reversing.
	// It will iterates entire input string.
	for fast := 0; fast < len(byteStr); fast++ {
		// 1. Convert the character if it's not blank.
		if byteStr[fast] != ' ' {
			// 3. Add a single blank to separate each word.
			if slow != 0 {
				byteStr[slow] = ' '
				slow++
			}
			// 2. Store current character in the slice of slow index if it's not blank.
			for fast < len(byteStr) && byteStr[fast] != ' ' {
				byteStr[slow] = byteStr[fast]
				slow++
				fast++
			}
		}
	}
	// Reverse the characters in string.
	ReverseStringApproach2(byteStr[0:slow])

	start := 0
	len := len(byteStr[0:slow])
	// Reverse each word in string.
	for i := 0; i <= len; i++ {
		// Remember to calculate the ending index of last word.
		// Avoid index out of range issue, put 'i == len' in front of byteStr[i] == ' '.
		if i == len || byteStr[i] == ' ' {
			ReverseStringApproach2(byteStr[start:i])
			start = i + 1
		}
	}
	return string(byteStr[0:slow])
}
