package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var reversed *ListNode

	for head != nil {
		head.Next, head, reversed = reversed, head.Next, head
	}

	return reversed
}
