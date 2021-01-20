package removeduplicatesfromsortedarray

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	input := []int{1, 1, 2}
	except := 2
	exceptAfter := []int{1, 2}
	actaul := removeDuplicates(input)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %v, Output: %d, But: %d", input, except, actaul)
	}
	if !reflect.DeepEqual(exceptAfter, input[:except]) {
		t.Errorf("Input: %v should become %v", input[:except], exceptAfter)
	}
}

func Test2(t *testing.T) {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	except := 5
	exceptAfter := []int{0, 1, 2, 3, 4}
	actaul := removeDuplicates(input)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %v, Output: %d, But: %d", input, except, actaul)
	}
	if !reflect.DeepEqual(exceptAfter, input[:except]) {
		t.Errorf("Input: %v should become %v", input[:except], exceptAfter)
	}
}
