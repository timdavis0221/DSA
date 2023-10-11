package str

// ReverseString2 reverses the first k characters for every 2k characters counting from the start
// of the string. If there are fewer than k characters left, reverse all of them. If there are less
// than 2k but greater than or equal to k characters, then reverse the first k characters and leave
// the other original.
//
// Constraints:
// String consists of only lowercase English letters.
// 1 <= k <= 10^4.
func ReverseString2(s string, k int) string {
	strByteSlice := []byte(s)
	strLen := len(s)
	// Use k as the end pointer.
	end := k
	// Use index i as the start pointer.
	// Reverse the string for every 2k characters.
	for i := 0; i < strLen; i += (2 * k) {
		// Reverse the first k characters (k <= length) if the range of the string starting from the
		// start pointer and stopping at the end pointer is less than or equal to the length of string.
		if (i + end) <= strLen {
			ReverseStringApproach2(strByteSlice[i : i+end])
		} else {
			// Reverse all of the characters. (k > length)
			ReverseStringApproach2(strByteSlice[i:])
		}
	}
	return string(strByteSlice)
}
