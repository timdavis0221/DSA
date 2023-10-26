package array

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
	// The fast pointer indicates the end position of the current range of subarray.
	// The slow pointer indicates the start position of the current range of subarray.
	slow, fast, sum := 0, 0, 0
	len := len(nums)
	// Default minimum length for non-existent cases.
	defaultMinLen := len + 1
	// The number of outer loop depends on the length of input size of nums.
	for fast < len {
		sum += nums[fast]
		// The inner loop only moves the slow pointer when the sum is greater than or equal to target.
		// Also, each elements are operated in two times at most.
		for sum >= target {
			minLen := fast - slow + 1
			if minLen < defaultMinLen {
				defaultMinLen = minLen
			}
			sum -= nums[slow]
			// Move the start position forward.
			slow++
		}
		fast++
	}
	// There isn't any matched minimum length after going through the iteration.
	// So indicate the length of input array plus 1 as the condition for non-existent cases.
	if fast+1 == defaultMinLen {
		return 0
	}
	return defaultMinLen
}
