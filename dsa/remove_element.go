package dsa

// RemoveElement returns the number of elements in nums which are not equal to val
// with two pointer method.
//
// Example 1:
// Input = [3,2,2,3], val = 3, output = 2
//
// Example 2:
// Input = [0,1,2,2,3,0,4,2], val = 2, output = 5
func RemoveElement(nums []int, val int) int {
	// Slow pointer for calculating the final nums of elements.
	count := 0
	// Use index of loop as fast pointer to iterate all elements.
	for index := range nums {
		if nums[index] != val {
			// Check the final elements of slice.
			nums[count] = nums[index]
			count++
		}
	}
	return count
}
