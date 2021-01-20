package forsum

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	input := []int{1, 0, -1, 0, -2, 2}
	target := 0
	except := [][]int{[]int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}}
	actaul := fourSum(input, target)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %v, Target: %d, Output:%v, But: %v", input, target, except, actaul)
	}
}

func Test2(t *testing.T) {
	input := []int{0, 0, 0, 0}
	target := 0
	except := [][]int{[]int{0, 0, 0, 0}}
	actaul := fourSum(input, target)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %v, Target: %d, Output:%v, But: %v", input, target, except, actaul)
	}
}

func Test3(t *testing.T) {
	input := []int{-2, 0, 0, 3, 3, -1}
	target := 5
	except := [][]int{[]int{-1, 0, 3, 3}}
	actaul := fourSum(input, target)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %v, Target: %d, Output:%v, But: %v", input, target, except, actaul)
	}
}
