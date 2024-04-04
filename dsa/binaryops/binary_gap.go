package binaryops

import "fmt"

// Binary Gap finds and returns the longest distance between 1's in the binary of input number.
func BinaryGap(n int) int {

	binary := fmt.Sprintf("%b", n)

	indices := []int{}
	for i, v := range binary {
		if v == rune('1') {
			indices = append(indices, i)
		}
	}
	// There are no any distances in the binary representation.
	if len(indices) == 1 {
		return 0
	}

	maxGap := 0
	for i := 0; i < len(indices)-1; i++ {
		maxGap = max(maxGap, indices[i+1]-indices[i])
	}
	return maxGap
}
