package dsa

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

// TwoSumApproach3 returns a indices of the two numbers.
func TwoSumApproach3(nums []int, target int) []int {

	// fmt.Println("Run TwoSumApproach3")

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

// MaxProfit returns a max profit from a give array.
func MaxProfit(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit < 0 {
				continue
			}
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}
