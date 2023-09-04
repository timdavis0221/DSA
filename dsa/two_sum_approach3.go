package dsa

// TwoSumApproach3 returns indices of the two numbers such that they add up to target.
//
// Example 1:
// Input = [2,7,11,15], target = 9, output = [0,1]
//
// Example 2:
// Input = [3,2,4], target = 6, output = [1,2]
func TwoSumApproach3(nums []int, target int) []int {
	// Hash table to store the target indies.
	m := make(map[int]int)

	for index, num := range nums {
		complement := target - num
		// Return index of complement (num) and current index if complement (num) had been kept before.
		if v, found := m[complement]; found {
			return []int{v, index}
		}
		// Keep current number as hash table's key and its corresponding value.
		// This index of value must be shown as complement in the following iteration.
		m[num] = index
	}
	return nil
}
