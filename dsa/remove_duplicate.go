package dsa

func RemoveDuplicate(nums []int) int {
	// [1, 1, 2]
	// [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
	current := 0
	for i := 1; i < len(nums); i++ {
		if nums[current] == nums[i] {
			continue
		}
		// if nums[current] != nums[i] {
		current++
		nums[current] = nums[i]
		// }
	}
	return current + 1
}
