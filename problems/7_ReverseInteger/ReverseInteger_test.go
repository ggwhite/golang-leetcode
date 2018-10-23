package reverseinteger

import (
	"testing"
)

func Test1(t *testing.T) {
	input := 123
	except := 321
	actaul := reverse(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%d, But: %d", input, except, actaul)
	}
}

func Test2(t *testing.T) {
	input := -123
	except := -321
	actaul := reverse(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%d, But: %d", input, except, actaul)
	}
}

func Test3(t *testing.T) {
	input := 120
	except := 21
	actaul := reverse(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%d, But: %d", input, except, actaul)
	}
}

func Test4(t *testing.T) {
	input := 1
	except := 1
	actaul := reverse(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%d, But: %d", input, except, actaul)
	}
}

func Test5(t *testing.T) {
	input := 10
	except := 1
	actaul := reverse(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%d, But: %d", input, except, actaul)
	}
}

func Test6(t *testing.T) {
	input := 2147483648
	except := 0
	actaul := reverse(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%d, But: %d", input, except, actaul)
	}
}

func Test9(t *testing.T) {
	input := -2147483649
	except := 0
	actaul := reverse(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%d, But: %d", input, except, actaul)
	}
}
