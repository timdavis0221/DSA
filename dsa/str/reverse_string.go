package str

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

// ReverseStringApproach3 makes the less code and more efficient by using the current length of the string to swap the characters.
func ReverseStringApproach3(s []byte) {
	currentLen := len(s)
	for i := 0; i < currentLen; i++ {
		s[i], s[currentLen-1] = s[currentLen-1], s[i]
		currentLen--
	}
}
