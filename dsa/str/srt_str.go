package str

// StrStr returns the index of the first occurrence of needle in haystack, or -1 if needle is
// not part of haystack.
func StrStr(haystack string, needle string) int {

	lengthHaystack := len(haystack)
	lengthNeedle := len(needle)
	// The length of haystack must includes the length of needle if needle is part of the haystack.
	for i := 0; i <= lengthHaystack-lengthNeedle; i++ {
		if haystack[i:i+lengthNeedle] == needle {
			return i
		}
	}
	return -1
}
