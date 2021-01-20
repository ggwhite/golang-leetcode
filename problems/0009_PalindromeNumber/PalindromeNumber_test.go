package palindromenumber

import (
	"testing"
)

func Test1(t *testing.T) {
	input := 121
	except := true
	actaul := isPalindrome(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%t, But: %t", input, except, actaul)
	}
}

func Test2(t *testing.T) {
	input := -121
	except := false
	actaul := isPalindrome(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%t, But: %t", input, except, actaul)
	}
}

func Test3(t *testing.T) {
	input := 10
	except := false
	actaul := isPalindrome(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%t, But: %t", input, except, actaul)
	}
}

func Test4(t *testing.T) {
	input := 1234554321
	except := true
	actaul := isPalindrome(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%t, But: %t", input, except, actaul)
	}
}

func Test5(t *testing.T) {
	input := 1210
	except := false
	actaul := isPalindrome(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%t, But: %t", input, except, actaul)
	}
}

func Test6(t *testing.T) {
	input := 1
	except := true
	actaul := isPalindrome(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%t, But: %t", input, except, actaul)
	}
}

func Test7(t *testing.T) {
	input := 11
	except := true
	actaul := isPalindrome(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%t, But: %t", input, except, actaul)
	}
}

func Test8(t *testing.T) {
	input := 1221
	except := true
	actaul := isPalindrome(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%t, But: %t", input, except, actaul)
	}
}
