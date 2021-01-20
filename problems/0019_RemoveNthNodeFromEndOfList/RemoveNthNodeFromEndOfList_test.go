package removenthnodefromendoflist

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
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	target := 2
	except := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}
	actaul := removeNthFromEnd(input, target)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %v, Target: %d, Output:%v, But: %v", input, target, except, actaul)
	}
}

func Test2(t *testing.T) {
	input := &ListNode{
		Val:  1,
		Next: nil,
	}
	target := 1
	actaul := removeNthFromEnd(input, target)
	if actaul != nil {
		t.Errorf("Input: %v, Target: %d, Output:%v, But: %v", input, target, nil, actaul)
	}
}

func Test3(t *testing.T) {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	target := 1
	except := &ListNode{
		Val:  1,
		Next: nil,
	}
	actaul := removeNthFromEnd(input, target)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %v, Target: %d, Output:%v, But: %v", input, target, except, actaul)
	}
}

func Test4(t *testing.T) {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	target := 2
	except := &ListNode{
		Val:  2,
		Next: nil,
	}
	actaul := removeNthFromEnd(input, target)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %v, Target: %d, Output:%v, But: %v", input, target, except, actaul)
	}
}
