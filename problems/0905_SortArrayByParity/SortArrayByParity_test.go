package sortarraybyparity

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	input := []int{3, 1, 2, 4}
	expects := [][]int{
		[]int{2, 4, 3, 1},
		[]int{4, 2, 3, 1},
		[]int{2, 4, 1, 3},
		[]int{4, 2, 1, 3},
	}
	actual := sortArrayByParity(input)

	pass := false

	for _, expect := range expects {
		if reflect.DeepEqual(expect, actual) {
			pass = true
			break
		}
	}

	if !pass {
		t.Errorf("Input: %v, Output: %v, But: %v", input, expects, actual)
	}
}
