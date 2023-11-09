package linkedlist

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// RemoveElements removes all the nodes of the linked list that Node.val is equal to val
// and returns the new head (list).
func RemoveElements(head *ListNode, val int) *ListNode {
	// Remove the head node where the value is equal to the val.
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return head
	}
	// In this step, we ensure the following head.val is not equal to the val.
	// Set up two pointers that can track the position and address of the node during the
	// following node traversal.
	previous := head
	current := head.Next
	// Move two pointers previous and current until current's next is nil.
	for current != nil {
		// Drop current node by changing the previous node's link to the current node's next address
		// if current value is equal to val. Otherwise, move previous forward to current.
		if current.Val == val {
			previous.Next = current.Next
		} else {
			previous = current
		}
		// In the end of traversal, move current address to the next node.
		current = current.Next
	}
	return head
}

// CreateLinkedList generates a linked list based on the given int array.
func CreateLinkedList(values []int) *ListNode {
	// The fake head that points to a real head of linked list.
	dummyHead := &ListNode{}
	current := dummyHead
	for _, val := range values {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	PrintList(dummyHead.Next)
	fmt.Println()
	return dummyHead.Next
}

// PrintList prints all the values of the ListNode.
func PrintList(node *ListNode) {
	current := node
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
}
