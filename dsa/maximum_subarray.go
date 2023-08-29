package dsa

// MaximumSubarray finds the subarray with the largest sum and returns its sum.
func MaximumSubarray(nums []int) int {
	sum, positiveCounter, negativeCounter, negativeNum := 0, 0, 0, -10000
	len := len(nums)
	// For only 1 element case.
	if len == 1 {
		return nums[0]
	}
	// For all positive and negative number cases.
	for _, num := range nums {
		if num >= 0 {
			sum += num
			positiveCounter++
		} else {
			negativeCounter++
			if num > negativeNum {
				negativeNum = num
			}
		}
	}
	if positiveCounter == len-1 {
		return sum
	} else if negativeNum == len-1 {
		return negativeNum
	}
	// For normal cases.
	maxSum := nums[0]
	// subarray := make([]int, 0)

	for index := 1; index < len; index++ {
		currentNum := nums[index]
		currentSum := currentNum + nums[index-1] 
		if currentNum > currentSum {
			maxSum = currentNum
			// also make a subarray begin with this num.
		}
		// [-2, 2, 5, -11, 6]
		// Which one is the max sum now?
		
		// 1. Is current sum greater than max sum?
		// 2. Is current num greater than max sum?
		// >> Update the max sum with bigger one. >> max(currentNum, currentSum)
		// >> or keep current max sum. (do nothing)


		// 3. Is current num greater than current sum?
		// >> start a new subarray begin with current num.
		// >> or keep current subarray. (do nothing)



	}





	return maxSum
}