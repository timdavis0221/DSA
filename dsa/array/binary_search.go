package array

// Examples:
// nums = [-1,0,3,5,9,12], target = 9, output = 4
// nums = [-1,0,3,5,9,12], target = 2, output = 0
// nums = [5], target = 5, output = 0
//
// BinarySearch returns index of target in the given array which is sorted in ascending order.
func BinarySearch(nums []int, target int) int {
	startIndex := 0
	endIndex := len(nums) - 1
	// It's necessary to calculate the middle index when start index is equal to end index
	// because it might be the answer in final iteration.
	for startIndex <= endIndex {
		midIndex := (startIndex + endIndex) / 2
		// Return target's index or shrink the given array in each iteration.
		if nums[midIndex] > target {
			endIndex = midIndex - 1
		} else if nums[midIndex] < target {
			startIndex = midIndex + 1
		} else {
			return midIndex
		}
	}
	return -1
}
