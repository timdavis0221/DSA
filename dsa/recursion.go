package dsa

// Recursion practice
func ReturnNumOfArray(a []int) int {
	if len(a) == 0 {
		return 0
	}
	return 1 + ReturnNumOfArray(a[1:])
}
