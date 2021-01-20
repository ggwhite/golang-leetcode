package removeelement

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	input := []int{3, 2, 2, 3}
	val := 3
	except := 2
	exceptAfter := []int{2, 2}
	actaul := removeElement(input, val)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %v, Output: %d, But: %d", input, except, actaul)
	}
	if !reflect.DeepEqual(exceptAfter, input[:except]) {
		t.Errorf("Input: %v should become %v", input[:except], exceptAfter)
	}
}

func Test2(t *testing.T) {
	input := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	except := 5
	exceptAfter := []int{0, 1, 3, 0, 4}
	actaul := removeElement(input, val)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %v, Output: %d, But: %d", input, except, actaul)
	}
	if !reflect.DeepEqual(exceptAfter, input[:except]) {
		t.Errorf("Input: %v should become %v", input[:except], exceptAfter)
	}
}

func Test3(t *testing.T) {
	input := []int{1}
	val := 1
	except := 0
	exceptAfter := []int{}
	actaul := removeElement(input, val)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %v, Output: %d, But: %d", input, except, actaul)
	}
	if !reflect.DeepEqual(exceptAfter, input[:except]) {
		t.Errorf("Input: %v should become %v", input[:except], exceptAfter)
	}
}
