package dsa

// MinSubArrayLen returns the min length of a subarray whose sum is greater than
// or equal to target.
//
// Example 1:
// Input: [2,3,1,2,4,3], Target: 7, Output: 2 ([4,3])
// Example 2:
// Input: [1,4,4], Target: 4, Output: 1 ([4])
// Example 3:
// Input: [1,1,1,1,1,1,1,1], Target: 11, Output: 0
func MinSubArrayLen(nums []int, target int) int {
	p1, p2, sum := 0, 0, 0
	len := len(nums)
	// Default minimum length for non-existent cases.
	defaultMinLen := len + 1
	// This loop for moving pointer 2 until its length < length of input array in O(n) time.
	for p2 < len {
		sum += nums[p2]
		// This loop only moves pointer 1 when the sum is greater than target.
		// So, each elements are operated in two times at most.
		for sum >= target {
			minLen := p2 - p1 + 1
			if minLen < defaultMinLen {
				defaultMinLen = minLen
			}
			sum -= nums[p1]
			p1++
		}
		p2++
	}
	// There isn't any matched minimum length after going through the iteration.
	// So indicate the length of input array plus 1 as the condition for non-existent cases.
	if p2+1 == defaultMinLen {
		return 0
	}
	return defaultMinLen
}
