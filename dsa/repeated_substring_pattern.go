package dsa

import (
	"strings"
)

// RepeatedSubstringPattern checks whether the input string can be constructed by taking
// a substring of it and appending multiple copies of the substring together.
func RepeatedSubstringPattern(s string) bool {
	doubled := s + s
	// The input string contains repeated substring pattern if there is at least one input string
	// can be found in the doubled string.
	return strings.Contains(doubled[1:len(doubled)-1], s)
}

// RepeatedSubstringPatternApproach2 checks whether the input string can be constructed by taking
// a substring of it and appending multiple copies of the substring together with KMP algorithm.
func RepeatedSubstringPatternApproach2(s string) bool {
	strLen := len(s)
	pt := prefixTable(s, strLen)
	// The first condition means there is a minimum repeated substring pattern occurs in the given string.
	// The second condition means the given string can be constructed by appending multiple copies of the
	// minimum repeated substring.
	if pt[strLen-1] != 0 && strLen%(strLen-pt[strLen-1]) == 0 {
		return true
	}
	return false
}

// prefixTable calculates the length of Longest Prefix Suffix from the given string and stores it in
// the last position of the returned slice.
func prefixTable(s string, len int) []int {
	// A slice that is used to store the length of the Longest Prefix Suffix (LPS).
	lps := make([]int, len)
	// Variable j is the slow pointer used to record the length of LPS and be the position of base value
	// that is used to compare with the next character in the input string.
	j := 0
	// The first value was set up as zero because there's not any LPS.
	lps[0] = j
	// Index i is the faster pointer used to iterate the original string and compare the current character
	// with base value. Also, it used to store the current length of LPS if LPS occurs.
	for i := 1; i < len; i++ {
		// There's not LPS occurs, grab the previous length of LPS as current length of LPS (zero).
		for j > 0 && s[j] != s[i] {
			// Set as zero directly?
			j = lps[j-1]
		}
		// There is a LPS occurs, increment j as the current length of LPS.
		if s[j] == s[i] {
			j++
		}
		// Update the length of LPS after each calculation.
		lps[i] = j
	}
	return lps
}
