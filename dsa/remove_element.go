package dsa

// RemoveElement...
func RemoveElement(nums []int, val int) int {
	// [3,2,2,3], val = 3
	// [0,1,2,2,3,0,4,2], val = 2
	count := 0
	for _, v := range nums {
		if v != val {
			nums[count] = v
			count++
		}
	}
	return count
}
