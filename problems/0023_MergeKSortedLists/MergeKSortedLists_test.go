package mergeksortedlists

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	input := []*ListNode{
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  6,
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
							Val: 4,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val:  6,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	actaul := mergeKLists(input)
	if !reflect.DeepEqual(except.String(), actaul.String()) {
		t.Errorf("Input: %v, Output:%s, But: %s", input, except.String(), actaul.String())
	}
}
