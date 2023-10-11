package array

// SortedSquares makes an O(n) improvements with Two Pointers.
func SortedSquares(nums []int) []int {
	// These two pointers start and end will traverse the given array start from left and right,
	// compare which square number is bigger then put it in a new slice in the same order.
	len := len(nums)
	start, end, result := 0, len-1, len-1
	ansSlice := make([]int, len)
	for start <= end {
		startValue, endValue := nums[start]*nums[start], nums[end]*nums[end]

		if startValue < endValue {
			ansSlice[result] = endValue
			end--
		} else {
			ansSlice[result] = startValue
			start++
		}
		result--
	}
	return ansSlice
}
