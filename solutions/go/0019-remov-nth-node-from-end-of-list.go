package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nth := &ListNode{
		Next: head,
	}
	newHead := nth

	for len, temp := 1, head.Next; temp != nil; len, temp = len+1, temp.Next {
		if len < n {
			continue
		}
		nth = nth.Next
	}
	nth.Next = nth.Next.Next

	return newHead.Next
}

// Another solution
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	nth := head
	len := 1

	for temp := head.Next; temp != nil; len, temp = len+1, temp.Next {
		if len <= n {
			continue
		}
		nth = nth.Next
	}

	if len == n {
		head = head.Next
	} else if nth.Next != nil {
		nth.Next = nth.Next.Next
	}

	return head
}
