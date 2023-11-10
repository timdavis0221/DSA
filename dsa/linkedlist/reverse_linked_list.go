package linkedlist

import "fmt"

// ReverseLinkedList reverses the given linked list iteratively.
func ReverseLinkedList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	fmt.Print("Reversed Linked List: ")
	PrintList(pre)
	return pre
}

// ReverseLinkedListRecursively reverses the given linked list recursively.
func ReverseLinkedListRecursively(head *ListNode) *ListNode {
	return reverse(nil, head)
}

func reverse(pre, head *ListNode) *ListNode {
	// Base case.
	cur := head
	if cur == nil {
		return pre
	}
	// Recursive case.
	// Keep the address of the next node before reversing.
	tmp := cur.next
	cur.next = pre
	// Moving previous and current pointer forward as sliding window to update the position of nodes.
	pre = cur
	cur = tmp

	PrintList(pre)
	return reverse(pre, cur)
}
