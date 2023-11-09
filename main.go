package main

import (
	"personalfile.app/yao/project_go/dsa/linkedlist"
)

func main() {
	// dsa.AppendSlices()
	// fmt.Println(dsa.ReturnNumOfArray([]int{3, 5, 6, 7}))
	// concurrency.Goroutine()

	/*
	 * Topic: Array
	 */
	// array.RemoveDuplicate([]int{1, 1, 2})
	// array.RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	// array.BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 9)
	// array.MaxProfit([]int{7, 1, 5, 3, 6, 4})
	// array.MaxProfitSolution2([]int{7, 1, 5, 3, 6, 4})
	// array.MaxProfitSolution3([]int{7, 1, 5, 3, 6, 4})
	// array.SortedSquares([]int{-4, -1, 0, 3, 10})
	// array.MinSubArrayLen([]int{1, 1, 1}, 4)
	// array.GenerateMatrix(3)
	// array.ProductExceptSelf([]int{1, 2, 3, 4})
	// array.ProductExceptSelfApproach2([]int{1, 2, 3, 4})
	// array.MaximumSubarray([]int{-2, 2, 5, -11, 6})
	// array.MaximumSubarrayApproach2([]int{-2, 2, 5, -11, 6})

	/*
	 * Topic: Sorting
	 */
	// sorting.SortedSquares([]int{-4, -1, 0, 3, 10})

	/*
	 * Topic: Searching
	 */
	// searching.SearchInRotatedSortedArray([]int{4, 5, 6, 7, 0, 1, 2}, 0)

	/*
	 * Topic: String
	 */
	// str.ReverseString([]byte{0x48, 0x65, 0x6c, 0x6c, 0x6f})
	// str.ReverseString2("hello", 2)
	// str.ValidAnagram("race", "care")
	// str.ValidAnagramApproach2("race", "care")
	// str.ValidAnagramApproach3("race", "care")
	// str.ValidPalindrome("!043XjqjX043!")
	// str.ReverseWords("the    y blue is sky")
	// str.ReverseWords2("     the y blue is sky")
	// str.RepeatedSubstringPattern("aba")
	// str.RepeatedSubstringPatternApproach2("aabbccaabbcc")
	// str.StrStr("sadbutsad", "sad")

	/*
	 * Topic: Linked List
	 */
	head := linkedlist.CreateLinkedList([]int{1, 2, 6, 3, 4, 5, 6})
	linkedlist.PrintList(linkedlist.RemoveElements(head, 6))

	/*
	 * Topic: Hash Table
	 */
	// hashtable.TwoSumApproach3([]int{2, 3, 5, 7}, 9)
}
