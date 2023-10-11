package array

// RemoveDuplicate returns the numbers of the first k elements that contains the unique elements
// in the order they were present in the given array nums initially using two pointer method.
//
// Example 1:
// Input = [1, 1 ,2], output = 2 (nums = [1, 2, _])
//
// Example 2:
// Input = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4], output = 5 (nums = [0, 1, 2, 3, 4, _, _, _, _, _])
func RemoveDuplicate(nums []int) int {
	// Slow pointer to record current element.
	current := 0
	// Fast pointer to iterate all elements to do comparison.
	// If current element is equal to the next element, skip this iteration. (Update fast pointer)
	// If current element isn't equal to the next element, increase slow pointer to update current
	// element with the next element.
	for i := 1; i < len(nums); i++ {
		if nums[current] == nums[i] {
			continue
		}
		current++
		nums[current] = nums[i]
	}
	return current + 1
}
