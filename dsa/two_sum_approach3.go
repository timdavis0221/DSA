package dsa

import "fmt"

// TwoSumApproach3 returns a indices of the two numbers.
func TwoSumApproach3(nums []int, target int) []int {

	fmt.Println("Run TwoSumApproach3")

	m := make(map[int]int)

	for index, num := range nums {
		complement := target - num
		if v, found := m[complement]; found {
			return []int{v, index}
		}
		m[num] = index
	}
	return nil
}
