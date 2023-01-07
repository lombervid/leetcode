package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	list := &ListNode{}
	last := list

	for l1 != nil || l2 != nil || carry > 0 {
		var num1, num2 int

		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}

		sum := num1 + num2 + carry
		carry = 0

		if sum >= 10 {
			sum -= 10
			carry = 1
		}

		last.Next = &ListNode{Val: sum}
		last = last.Next
	}

	return list.Next
}
