package searching

// SearchInRotatedSortedArray searches the index of target number in the given rotated sorted
// array in O(log n) time.
func SearchInRotatedSortedArray(nums []int, target int) int {
	// Use left and right pointer to search the input array.
	start, end := 0, len(nums)-1
	for start <= end {
		pivot := (start + end) / 2
		if nums[pivot] == target {
			return pivot
		}
		// Key point is to check whether the left or right halves is sorted because the
		// input ascending array has been rotated sorted, adopt binary search directly
		// will lead the false results.

		// The left halves must be rotated sorted if mid value is greater than or equal to start value.
		// Otherwise, the right halves is original ascending sorted.
		// In order to use binary search, we must recognize which space is sorted.
		if nums[start] <= nums[pivot] {
			// Target is in the range of left halves.
			if nums[start] <= target && target < nums[pivot] {
				end = pivot - 1
			} else {
				start = pivot + 1
			}
		} else {
			// Target is in the range of right halves.
			if nums[pivot] < target && target <= nums[end] {
				start = pivot + 1
			} else {
				end = pivot - 1
			}
		}
	}
	return -1
}
