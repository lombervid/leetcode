package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	middle := head
	head = head.Next

	for i := 0; head != nil; i, head = i+1, head.Next {
		if i%2 == 0 {
			middle = middle.Next
		}
	}

	return middle
}
