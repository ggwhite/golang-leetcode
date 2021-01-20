package powxn

import (
	"testing"
)

func Test1(t *testing.T) {
	x := 2.00000
	n := 10
	except := 1024.00000
	actaul := myPow(x, n)
	if except != actaul {
		t.Errorf("Input:%f, %d, Output: %f, But: %f", x, n, except, actaul)
	}
}

func Test2(t *testing.T) {
	x := 2.10000
	n := 3
	except := 9.261000
	actaul := myPow(x, n)
	if except != actaul {
		t.Errorf("Input:%f, %d, Output: %f, But: %f", x, n, except, actaul)
	}
}

func Test3(t *testing.T) {
	x := 2.00000
	n := -2
	except := 0.25000
	actaul := myPow(x, n)
	if except != actaul {
		t.Errorf("Input:%f, %d, Output: %f, But: %f", x, n, except, actaul)
	}
}
