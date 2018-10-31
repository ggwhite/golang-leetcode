package mergetwosortedlists

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	input1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	input2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	except := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
	}
	actaul := mergeTwoLists(input1, input2)
	if !reflect.DeepEqual(except.String(), actaul.String()) {
		t.Errorf("Input: %s, %s, Output:%s, But: %s", input1.String(), input2.String(), except.String(), actaul.String())
	}
}
