package containerwithmostwater

import (
	"testing"
)

func Test1(t *testing.T) {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	except := 49
	actaul := maxArea(input)
	if except != actaul {
		t.Errorf("Input: %v, Output:%d, But: %d", input, except, actaul)
	}
}

func Test2(t *testing.T) {
	input := []int{5, 2, 4, 20, 7, 3, 9, 5, 10, 1}
	except := 50
	actaul := maxArea(input)
	if except != actaul {
		t.Errorf("Input: %v, Output:%d, But: %d", input, except, actaul)
	}
}
