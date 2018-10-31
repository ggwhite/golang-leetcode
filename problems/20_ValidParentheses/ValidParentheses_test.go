package validparentheses

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	input := "()"
	except := true
	actaul := isValid(input)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %s, Output:%t, But: %t", input, except, actaul)
	}
}

func Test2(t *testing.T) {
	input := "()[]{}"
	except := true
	actaul := isValid(input)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %s, Output:%t, But: %t", input, except, actaul)
	}
}

func Test3(t *testing.T) {
	input := "(]"
	except := false
	actaul := isValid(input)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %s, Output:%t, But: %t", input, except, actaul)
	}
}

func Test4(t *testing.T) {
	input := "([)]"
	except := false
	actaul := isValid(input)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %s, Output:%t, But: %t", input, except, actaul)
	}
}

func Test5(t *testing.T) {
	input := "{[]}"
	except := true
	actaul := isValid(input)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %s, Output:%t, But: %t", input, except, actaul)
	}
}

func Test6(t *testing.T) {
	input := "{[ab]}"
	except := false
	actaul := isValid(input)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %s, Output:%t, But: %t", input, except, actaul)
	}
}

func Test7(t *testing.T) {
	input := "{[]"
	except := false
	actaul := isValid(input)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %s, Output:%t, But: %t", input, except, actaul)
	}
}

func Test8(t *testing.T) {
	input := "{}{}()[]"
	except := true
	actaul := isValid(input)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %s, Output:%t, But: %t", input, except, actaul)
	}
}

func Test9(t *testing.T) {
	input := "(([]){})"
	except := true
	actaul := isValid(input)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %s, Output:%t, But: %t", input, except, actaul)
	}
}

func Test10(t *testing.T) {
	input := "{}{{}}"
	except := true
	actaul := isValid(input)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %s, Output:%t, But: %t", input, except, actaul)
	}
}
