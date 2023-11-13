package linkedlist

// HasCycle checks if the given linked list has a cycle in it.
func HasCycle(head *ListNode) bool {
	visitedNodes := make(map[*ListNode]bool)
	current := head
	for current.next != nil {
		if visitedNodes[current] {
			return true
		}
		visitedNodes[current] = true
		current = current.next
	}
	return false
}

// HasCycleApproach2 detects whether the given linked list has a cycle by
// Floyd's Tortoise and Hare Algorithm.
func HasCycleApproach2(head *ListNode) bool {
	// Define the faster pointer that moves two steps in each iteration and assume head
	// as the slow pointer that moves one step in the same iteration.
	fast := head
	// Loop until reach the end of the linked list.
	for fast != nil && fast.next != nil {
		head = head.next
		fast = fast.next.next
		// According to Floyd's Tortoise and Hare Algorithm, Tortoise and Hare would meet
		// at the same position if there's a cycle in the path, which means the memory address
		// of their position would be the same, so if these two position are the same, there
		// must be a cycle in the linked list.
		if fast == head {
			return true
		}
	}
	return false
}

// SinglyLinkedList generates a singly linked list.
func SinglyLinkedList(values []int) *ListNode {
	head := &ListNode{}
	current := head
	for index, val := range values {
		current.val = val
		if index != len(values)-1 {
			current.next = &ListNode{}
		}
		current = current.next
	}
	// Do not print the linked list with a cycle that would lead to an infinite loop.
	PrintList(head)
	return head
}

// CycledLinkedList generates a linked list with a cycle.
func CycledLinkedList(values []int) *ListNode {
	head := &ListNode{}
	current := head
	for index, val := range values {
		current.val = val
		current.next = &ListNode{}
		current = current.next
		if index == len(values)-1 {
			current.next = head.next
		}
	}
	return head
}
