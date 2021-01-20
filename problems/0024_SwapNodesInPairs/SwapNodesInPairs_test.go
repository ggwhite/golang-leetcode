package swapnodesinpairs

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	except := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	actaul := swapPairs(input)
	if !reflect.DeepEqual(except.String(), actaul.String()) {
		t.Errorf("Input: %s, Output: %s, But: %s", input.String(), except.String(), actaul.String())
	}
}

func Test2(t *testing.T) {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	except := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	actaul := swapPairs(input)
	if !reflect.DeepEqual(except.String(), actaul.String()) {
		t.Errorf("Input: %s, Output: %s, But: %s", input.String(), except.String(), actaul.String())
	}
}

func Test3(t *testing.T) {
	input := &ListNode{
		Val:  1,
		Next: nil,
	}
	except := &ListNode{
		Val:  1,
		Next: nil,
	}
	actaul := swapPairs(input)
	if !reflect.DeepEqual(except.String(), actaul.String()) {
		t.Errorf("Input: %s, Output: %s, But: %s", input.String(), except.String(), actaul.String())
	}
}

func Test4(t *testing.T) {
	actaul := swapPairs(nil)
	if actaul != nil {
		t.Errorf("Input: nil, Output: nil, But: %s", actaul.String())
	}
}
