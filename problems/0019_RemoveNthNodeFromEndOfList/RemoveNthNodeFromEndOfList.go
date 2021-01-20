// Package removenthnodefromendoflist https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
package removenthnodefromendoflist

import "strconv"

/*
Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.

Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?
*/

// ListNode is define on leetcode
type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	if n.Next != nil {
		return strconv.Itoa(n.Val) + "->" + n.Next.String()
	}
	return strconv.Itoa(n.Val)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	var nodes []*ListNode
	for ; head != nil; head = head.Next {
		nodes = append(nodes, head)
	}
	l := len(nodes)
	if l == 0 || (l == 1 && l == n) {
		return nil
	}
	if l-n-1 < 0 {
		return nodes[1]
	}
	nodes[l-n-1].Next = nodes[l-n].Next
	return nodes[0]
}
