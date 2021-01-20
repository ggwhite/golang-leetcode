package generateparentheses

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	input := 3
	except := []string{
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()",
	}
	actaul := generateParenthesis(input)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %d, Output:%v, But: %v", input, except, actaul)
	}
}
