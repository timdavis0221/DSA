package dsa

// ReverseString reverses a string, must doing this by modifying the input array in-place with
// O(1) extra memory.
func ReverseString(s []byte) {
	len := len(s)
	end := len - 1
	// Swap the characters until start index is equal to end index or length of end minus one.
	for start := range s {
		tmp := s[start]
		s[start] = s[end]
		s[end] = tmp
		end--
		if start == end || start == (end-1) {
			break
		}
	}
}

// ReverseStringApproach2 makes the improvements with concepts of Two Pointers and refines code.
func ReverseStringApproach2(s []byte) {
	len := len(s)
	start, end := 0, len-1
	for start < end {
		// Swap the elements with tuple assignment.
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
