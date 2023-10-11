package array

import (
	"math"

	"personalfile.app/yao/project_go/lib"
)

// MaximumSubarray finds the subarray with the largest sum and returns its sum
// with Kadane's Algorithm.
func MaximumSubarray(nums []int) int {
	// Define the first element as default maximum sum.
	// Can not use zero as default value because current sum and current num might be negative number.
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

// MaximumSubarrayApproach2 returns the maximum sum of subarray from the given array
// with Divide and Conquer algorithm in O(n log n) time.
func MaximumSubarrayApproach2(nums []int) int {
	return maxSubarray(nums, 0, len(nums)-1)
}

func maxSubarray(nums []int, left, right int) int {
	// Base case.
	if left == right {
		return nums[left]
	}
	// Recursive case.
	mid := (left + right) / 2
	maxSumLeft := math.MinInt
	// It's necessary to set default current sum as zero because we will add all elements up
	// in specific order.
	currentSumLeft := 0
	// We must iterate elements from mid to left and mid+1 to right to make sure the final sum of
	// subarray is the maximum sum we retrieved from all subarray.
	for index := mid; index >= left; index-- {
		currentSumLeft += nums[index]
		maxSumLeft = lib.Max(maxSumLeft, currentSumLeft)
	}
	maxSumRight := math.MinInt
	currentSumRight := 0
	for index := mid + 1; index <= right; index++ {
		currentSumRight += nums[index]
		maxSumRight = lib.Max(maxSumRight, currentSumRight)
	}
	// Calculate the left sum, right sum and left sum plus right sum to obtain the final max sum.
	subSumLeft := maxSubarray(nums, 0, mid)
	subSumRight := maxSubarray(nums, mid+1, right)
	maxSum := lib.Max(lib.Max(subSumLeft, subSumRight), maxSumLeft+maxSumRight)
	return maxSum
}
