package array

// ProductExceptSelf returns an int array answer so that answer[i] is equal to the product of all the
// elements of nums except nums[i].
//
// Constraint: write an algorithm that runs in O(n) time and without using the division operation.
//
// Example 1:
// Input: nums = [1,2,3,4], Output: [24,12,8,6]
// Example 2:
// Input: nums = [-1,1,0,-3,3], Output: [0,0,9,0,0]
func ProductExceptSelf(nums []int) []int {
	len := len(nums)
	res := make([]int, len)
	leftProducts := make([]int, len)
	rightProducts := make([]int, len)
	// nums:           1,  2,  3, 4
	// left product:   x,  1,  2, 6
	// right product:  24, 12, 4, x
	//
	// Result is multiply each element in the left product with corresponding element
	// in the right product. So we fill in these blank spaces x as 1 for calculation.
	leftProducts[0] = 1
	rightProducts[len-1] = 1

	for index := range nums[1:len] {
		index++
		leftProducts[index] = nums[index-1] * leftProducts[index-1]
	}
	for i := len - 2; i >= 0; i-- {
		rightProducts[i] = nums[i+1] * rightProducts[i+1]
	}
	for i := 0; i < len; i++ {
		res[i] = leftProducts[i] * rightProducts[i]
	}
	return res
}

// ProductExceptSelfApproach2 makes space optimization in O(1).
//
// Note: this is follow up challenge.
func ProductExceptSelfApproach2(nums []int) []int {
	len := len(nums)
	res := make([]int, len)
	// Replace the slice of left product with result slice to store all the elements in left product.
	// Also, remove extra slice of right product.
	res[0] = 1
	for index := range nums[1:len] {
		index++
		res[index] = nums[index-1] * res[index-1]
	}
	// The suffix stores the current value of suffix, move from right to left
	// for updating the result slice.
	suffix := 1
	for index := len - 1; index >= 0; index-- {
		// This formula is equal to leftProducts[i] * rightProducts[i] above.
		res[index] = res[index] * suffix
		// The next (previous) value of right product.
		suffix *= nums[index]
	}
	return res
}
