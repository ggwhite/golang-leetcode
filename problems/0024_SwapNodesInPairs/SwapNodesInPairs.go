// Package swapnodesinpairs https://leetcode.com/problems/swap-nodes-in-pairs/description/
package swapnodesinpairs

import "fmt"

/*
Given a linked list, swap every two adjacent nodes and return its head.

Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.
Note:

Your algorithm should use only constant extra space.
You may not modify the values in the list's nodes, only nodes itself may be changed.
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

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmp := head
	next := head.Next
	tmp.Next = swapPairs(next.Next)
	head = next
	head.Next = tmp

	return head
}
