package dsa

import (
	"personalfile.app/yao/project_go/lib"
)

// MaximumSubarray finds the subarray with the largest sum and returns its sum
// with Kadane's Algorithm.
func MaximumSubarray(nums []int) int {
	maxSum := nums[0]
	currentSum := maxSum
	for _, num := range nums[1:] {
		// Add up all elements until current element to gain the current sum.
		currentSum += num
		// The key point is we have to make sure that the current sum here is guaranteed to
		// the maximum sum for now, it can be zero or a negative number.
		currentSum = lib.Max(currentSum, num)
		maxSum = lib.Max(currentSum, maxSum)
	}
	return maxSum
}
