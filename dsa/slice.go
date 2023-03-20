package dsa

import "fmt"

func AppendSlices() {
	nums := make([]int, 0)
	fmt.Printf("init nums: %v, addr: %p\n", nums, &nums)

	nums = append(nums, 1)
	fmt.Printf("1 element in nums: %v, addr: %p, len: %d, cap: %d\n", nums, &nums[0], len(nums), cap(nums))
	nums = append(nums, 2)
	fmt.Printf("2 element in nums: %v, addr: %p, len: %d, cap: %d\n", nums, &nums[0], len(nums), cap(nums))
	nums = append(nums, 3)
	fmt.Printf("3 element in nums: %v, addr: %p, len: %d, cap: %d\n", nums, &nums[0], len(nums), cap(nums))
	nums = append(nums, 4)
	fmt.Printf("4 element in nums: %v, addr: %p, len: %d, cap: %d\n", nums, &nums[0], len(nums), cap(nums))
	nums = append(nums, 5)
	fmt.Printf("5 element in nums: %v, addr: %p, len: %d, cap: %d\n\n", nums, &nums[0], len(nums), cap(nums))

	// nums = [1, 2, 3, 4, 5] (old slice)
	// newNums = [1, 2, 3, 4, 5, 6] (updated slice since we still have sufficient cap 8)
	newNums := append(nums, 6)
	fmt.Printf("append 6th element in nums: %v, addr: %p, len: %d, cap: %d, self addr: %p\n", nums, &nums[0], len(nums), cap(nums), &nums)
	fmt.Printf("newNums: %v, addr: %p, len: %d, cap: %d, self addr: %p\n\n", newNums, &newNums[0], len(newNums), cap(newNums), &newNums)

	fmt.Println("Override 6th element in underlying array since we used old slice (nums) to update it")
	// nums = [1, 2, 3, 4, 5, 7]
	nums = append(nums, 7)
	fmt.Printf("append 7th element in nums: %v, addr: %p, len: %d, cap: %d, self addr: %p\n\n", nums, &nums[0], len(nums), cap(nums), &nums)

	newNums2 := append(newNums, 8)
	fmt.Printf("append 8th element in newNums2: %v, addr: %p, len: %d, cap: %d, self addr: %p\n", newNums2, &newNums2[0], len(newNums2), cap(newNums2), &newNums2)
	fmt.Printf("newNums: %v, addr: %p, len: %d, cap: %d, self addr: %p\n\n", newNums, &newNums[0], len(newNums), cap(newNums), &newNums)

	fmt.Println("Allocate independant array to each slice")
	slice1 := make([]int, 0)
	slice1 = append(slice1, 1)
	slice1 = append(slice1, 2)
	slice1 = append(slice1, 3)
	slice1 = append(slice1, 4)
	slice1 = append(slice1, 5)

	slice2 := make([]int, len(slice1), len(slice1)+1)
	minLen := copy(slice2, slice1)
	fmt.Printf("Minimum len of dst and src slice: %d\n", minLen)

	slice2 = append(slice2, 6)
	fmt.Printf("slice2: %v, addr of array: %p\n", slice2, &slice2[0])
	slice1 = append(slice1, 7)
	fmt.Printf("slice1: %v, addr of array: %p\n", slice1, &slice1[0])
}
