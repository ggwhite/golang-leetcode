// Package reversenodesinkgroup https://leetcode.com/problems/reverse-nodes-in-k-group/description/
package reversenodesinkgroup

import "fmt"

/*
Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

Example:

Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5

Note:

Only constant extra memory is allowed.
You may not alter the values in the list's nodes, only nodes itself may be changed.

*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	if n.Next != nil {
		return fmt.Sprintf("%d -> %s", n.Val, n.Next.String())
	}
	return fmt.Sprintf("%d", n.Val)
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 || k == 1 {
		return head
	}
	// find group start and end
	start, end := head, head
	for i := 0; i < k-1; i++ {
		end = end.Next
		if end == nil {
			return head
		}
	}
	if end.Next == nil {
		return revers(head)
	}

	tmp := end.Next
	end.Next = nil
	head = reverseKGroup(start, k)
	start, end = head, head
	for i := 0; i < k-1; i++ {
		end = end.Next
	}
	end.Next = reverseKGroup(tmp, k)
	return head
}

func revers(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var arr []*ListNode
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	n := len(arr)
	for i, j := 0, n-1; i < n/2; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	for i := 0; i < n-1; i++ {
		arr[i].Next = arr[i+1]
	}
	arr[n-1].Next = nil
	return arr[0]
}
