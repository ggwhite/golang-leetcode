// Package mergetwosortedlists https://leetcode.com/problems/merge-two-sorted-lists/description/
package mergetwosortedlists

import "strconv"

/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/
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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil && l2 != nil {
		return l2
	} else if l1 != nil && l2 == nil {
		return l1
	}

	var current *ListNode
	if l1.Val < l2.Val {
		current = l1
		current.Next = mergeTwoLists(l1.Next, l2)
	} else {
		current = l2
		current.Next = mergeTwoLists(l1, l2.Next)
	}

	return current
}
