package reversenodesinkgroup

import (
	"reflect"
	"testing"
)

// func Test1(t *testing.T) {
// 	input := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 3,
// 				Next: &ListNode{
// 					Val: 4,
// 					Next: &ListNode{
// 						Val:  5,
// 						Next: nil,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	k := 2
// 	except := &ListNode{
// 		Val: 2,
// 		Next: &ListNode{
// 			Val: 1,
// 			Next: &ListNode{
// 				Val: 4,
// 				Next: &ListNode{
// 					Val: 3,
// 					Next: &ListNode{
// 						Val:  5,
// 						Next: nil,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	actaul := reverseKGroup(input, k)
// 	if !reflect.DeepEqual(except.String(), actaul.String()) {
// 		t.Errorf("Input: %s, K: %d, Output: %s, But: %s", input.String(), k, except.String(), actaul.String())
// 	}
// }

// func Test2(t *testing.T) {
// 	input := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 3,
// 				Next: &ListNode{
// 					Val: 4,
// 					Next: &ListNode{
// 						Val:  5,
// 						Next: nil,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	k := 3
// 	except := &ListNode{
// 		Val: 3,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 1,
// 				Next: &ListNode{
// 					Val: 4,
// 					Next: &ListNode{
// 						Val:  5,
// 						Next: nil,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	actaul := reverseKGroup(input, k)
// 	if !reflect.DeepEqual(except.String(), actaul.String()) {
// 		t.Errorf("Input: %s, K: %d, Output: %s, But: %s", input.String(), k, except.String(), actaul.String())
// 	}
// }

// func Test3(t *testing.T) {
// 	input := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 3,
// 				Next: &ListNode{
// 					Val: 4,
// 					Next: &ListNode{
// 						Val:  5,
// 						Next: nil,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	k := 0
// 	except := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 3,
// 				Next: &ListNode{
// 					Val: 4,
// 					Next: &ListNode{
// 						Val:  5,
// 						Next: nil,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	actaul := reverseKGroup(input, k)
// 	if !reflect.DeepEqual(except.String(), actaul.String()) {
// 		t.Errorf("Input: %s, K: %d, Output: %s, But: %s", input.String(), k, except.String(), actaul.String())
// 	}
// }

// func Test4(t *testing.T) {
// 	input := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 3,
// 				Next: &ListNode{
// 					Val: 4,
// 					Next: &ListNode{
// 						Val:  5,
// 						Next: nil,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	k := 1
// 	except := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 3,
// 				Next: &ListNode{
// 					Val: 4,
// 					Next: &ListNode{
// 						Val:  5,
// 						Next: nil,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	actaul := reverseKGroup(input, k)
// 	if !reflect.DeepEqual(except.String(), actaul.String()) {
// 		t.Errorf("Input: %s, K: %d, Output: %s, But: %s", input.String(), k, except.String(), actaul.String())
// 	}
// }

// func Test5(t *testing.T) {
// 	actaul := reverseKGroup(nil, 3)
// 	if actaul != nil {
// 		t.Errorf("Input: nil, K: %d, Output: nil, But: %s", 3, actaul.String())
// 	}
// }

// func Test6(t *testing.T) {
// 	input := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 3,
// 				Next: &ListNode{
// 					Val: 4,
// 					Next: &ListNode{
// 						Val:  5,
// 						Next: nil,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	k := 6
// 	except := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 3,
// 				Next: &ListNode{
// 					Val: 4,
// 					Next: &ListNode{
// 						Val:  5,
// 						Next: nil,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	actaul := reverseKGroup(input, k)
// 	if !reflect.DeepEqual(except.String(), actaul.String()) {
// 		t.Errorf("Input: %s, K: %d, Output: %s, But: %s", input.String(), k, except.String(), actaul.String())
// 	}
// }

// func Test7(t *testing.T) {
// 	input := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 3,
// 				Next: &ListNode{
// 					Val: 4,
// 					Next: &ListNode{
// 						Val:  5,
// 						Next: nil,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	k := 5
// 	except := &ListNode{
// 		Val: 5,
// 		Next: &ListNode{
// 			Val: 4,
// 			Next: &ListNode{
// 				Val: 3,
// 				Next: &ListNode{
// 					Val: 2,
// 					Next: &ListNode{
// 						Val:  1,
// 						Next: nil,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	actaul := reverseKGroup(input, k)
// 	if !reflect.DeepEqual(except.String(), actaul.String()) {
// 		t.Errorf("Input: %s, K: %d, Output: %s, But: %s", input.String(), k, except.String(), actaul.String())
// 	}
// }

// func Test8(t *testing.T) {
// 	input := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 3,
// 				Next: &ListNode{
// 					Val: 4,
// 					Next: &ListNode{
// 						Val:  5,
// 						Next: nil,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	fmt.Println(revers(input))
// }

func Test9(t *testing.T) {
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
	k := 3
	except := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	actaul := reverseKGroup(input, k)
	if !reflect.DeepEqual(except.String(), actaul.String()) {
		t.Errorf("Input: %s, K: %d, Output: %s, But: %s", input.String(), k, except.String(), actaul.String())
	}
}
