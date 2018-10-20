package addtwonumbers

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	expect := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  8,
				Next: nil,
			},
		},
	}
	actual := addTwoNumbers(l1, l2)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("l1: %v, l2: %v, expect: %v, actual: %v", l1, l2, expect, actual)
	}
}

func Test2(t *testing.T) {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}

	l2 := &ListNode{
		Val:  2,
		Next: nil,
	}

	expect := &ListNode{
		Val:  3,
		Next: nil,
	}
	actual := addTwoNumbers(l1, l2)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("l1: %v, l2: %v, expect: %v, actual: %v", l1, l2, expect, actual)
	}
}

func Test3(t *testing.T) {
	l1 := &ListNode{
		Val:  8,
		Next: nil,
	}

	l2 := &ListNode{
		Val:  2,
		Next: nil,
	}

	expect := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	actual := addTwoNumbers(l1, l2)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("l1: %v, l2: %v, expect: %v, actual: %v", l1, l2, expect, actual)
	}
}
