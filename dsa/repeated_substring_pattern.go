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
