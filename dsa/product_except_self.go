package dsa

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
