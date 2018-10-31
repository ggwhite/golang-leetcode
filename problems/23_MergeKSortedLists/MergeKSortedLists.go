// Package mergeksortedlists https://leetcode.com/problems/merge-k-sorted-lists/description/
package mergeksortedlists

import "fmt"

/*
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
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

var mergetimes int

func mergeKLists(lists []*ListNode) *ListNode {
	// Sample 8ms
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	ans := merge(mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:]))
	fmt.Println(mergetimes)
	return ans

	/* My Solution 1 332ms
	var current *ListNode
	var minIdx int
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		if lists[minIdx] == nil || lists[i].Val < lists[minIdx].Val {
			minIdx = i
		}
	}
	if lists[minIdx] == nil {
		return nil
	}
	current = lists[minIdx]
	lists[minIdx] = lists[minIdx].Next
	current.Next = mergeKLists(lists)
	return current
	*/

	/* My Solution 2 204ms, add one by one
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	head := lists[0]
	for i := 1; i < n; i++ {
		head = merge(head, lists[i])
	}
	fmt.Println(mergetimes)
	return head
	*/
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	mergetimes++
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var current *ListNode
	if l1.Val < l2.Val {
		current = l1
		current.Next = merge(l1.Next, l2)
	} else {
		current = l2
		current.Next = merge(l1, l2.Next)
	}
	return current
}

/*
Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6

There are 9 nodes in input
split input as [1->4->5] [1->3->4, 2->6]
both that call mergeKLists and get two ListNode
then merge both node as one
*/
