package threesumclosest

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	input := []int{-1, 2, 1, -4}
	target := 1
	except := 2
	actaul := threeSumClosest(input, target)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %v, Target: %d, Output:%d, But: %d", input, target, except, actaul)
	}
}

func Test2(t *testing.T) {
	input := []int{-3, -3, -2, -2, -1, -1, -1, 0, 0, 0, 1, 1, 2, 3, 3, 4}
	target := 1
	except := 1
	actaul := threeSumClosest(input, target)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %v, Target: %d, Output:%d, But: %d", input, target, except, actaul)
	}
}

func Test3(t *testing.T) {
	input := []int{-3, -3, -2, -2, -1, -1, -1, 0, 0, 0, 1, 1, 2, 3, 3, 4}
	target := 20
	except := 10
	actaul := threeSumClosest(input, target)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %v, Target: %d, Output:%d, But: %d", input, target, except, actaul)
	}
}

func Test4(t *testing.T) {
	input := []int{-3, -3, -2, -2, -1, -1, -1, 0, 0, 0, 1, 1, 2, 3, 3, 4}
	target := -20
	except := -8
	actaul := threeSumClosest(input, target)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %v, Target: %d, Output:%d, But: %d", input, target, except, actaul)
	}
}

func Test5(t *testing.T) {
	input := []int{-100, -10, -10, -1, 2, 3, 4, 4, 4, 4, 5, 40, 70}
	target := -90
	except := -91
	actaul := threeSumClosest(input, target)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %v, Target: %d, Output:%d, But: %d", input, target, except, actaul)
	}
}

func Test6(t *testing.T) {
	input := []int{-100, -10, -1, 2, 3, 40, 70}
	target := 60
	except := 59
	actaul := threeSumClosest(input, target)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %v, Target: %d, Output:%d, But: %d", input, target, except, actaul)
	}
}

func Test7(t *testing.T) {
	input := []int{-100, -10, -1, 2, 3, 40, 70}
	target := 3
	except := 4
	actaul := threeSumClosest(input, target)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %v, Target: %d, Output:%d, But: %d", input, target, except, actaul)
	}
}
