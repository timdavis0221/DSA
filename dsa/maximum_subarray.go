package dsa

import (
	"personalfile.app/yao/project_go/lib"
)

// MaximumSubarray finds the subarray with the largest sum and returns its sum.
func MaximumSubarray(nums []int) int {
	maxSum := nums[0]
	currentSum := maxSum
	for _, num := range nums[1:] {
		currentSum += num
		maxSum = lib.Max(maxSum, currentSum)
		// If current number is greater than current sum and maximum sum, drop the
		// previous sum of subarray and set current number as maximum sum.
		// Also, use current number as the new sum of subarray.
		if num > currentSum && num > maxSum {
			maxSum = num
			currentSum = maxSum
		}
		// TODO: wrong answer in this case [3,-2,-3,-3,1,3,0].
		if currentSum < 0 && num < currentSum {
			currentSum = 0
		}
	}
	return maxSum
}
