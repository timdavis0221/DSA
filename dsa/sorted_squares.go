package dsa

import (
	"fmt"
	"math/rand"
)

// SortedSquaresSolution2 makes an O(n) improvements with Two Pointers.
func SortedSquaresSolution2(nums []int) []int {
	// These two pointers start and end will traverse the given array start from left and right,
	// compare which square number is bigger then put it in a new slice in the same order.
	len := len(nums)
	start, end, result := 0, len-1, len-1
	ansSlice := make([]int, len)
	for start <= end {
		startValue, endValue := nums[start]*nums[start], nums[end]*nums[end]

		if startValue < endValue {
			ansSlice[result] = endValue
			end--
		} else {
			ansSlice[result] = startValue
			start++
		}
		result--
	}
	return ansSlice
}

// SortedSquares returns an array of the squares of each number sorted in non-decreasing order.
//
// Example 1:
// Input: [-4,-1,0,3,10], output: [0,1,9,16,100]
// Example 2:
// Input: [-7,-3,2,3,11], output: [4,9,9,49,121]
func SortedSquares(nums []int) []int {

	for idx, num := range nums {
		nums[idx] = num * num
	}

	return quickSort(nums, 0, len(nums)-1)
}

// quickSort performs quick sorting algorithm in average time. O(n log n)
func quickSort(nums []int, start, end int) []int {
	var pivotIndex int
	less, mid, greater := make([]int, 0), make([]int, 0), make([]int, 0)
	// Base case.
	if len(nums) < 2 {
		return nums
	} else { // Partition the array with pivot index.
		// Random choose an index as pivot.
		pivotIndex = rand.Intn(end)
		fmt.Println("pivot index: ", pivotIndex)
		for _, num := range nums {
			fmt.Println("num: ", num)
			if num < nums[pivotIndex] {
				less = append(less, num)
				fmt.Println("less: ", less)
			} else if num > nums[pivotIndex] {
				greater = append(greater, num)
				fmt.Println("greater: ", greater)
			} else if num == nums[pivotIndex] {
				mid = append(mid, num)
			}
		}
	}
	// Combine the partitions.
	resultSlice := make([]int, 0)
	resultSlice = append(resultSlice, quickSort(less, 0, len(less)-1)...)
	resultSlice = append(resultSlice, mid...)
	resultSlice = append(resultSlice, quickSort(greater, 0, len(greater)-1)...)
	return resultSlice
}
