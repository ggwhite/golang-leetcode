package romantointeger

import (
	"testing"
)

func Test1(t *testing.T) {
	input := "III"
	except := 3
	actaul := romanToInt(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

func Test2(t *testing.T) {
	input := "IV"
	except := 4
	actaul := romanToInt(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

func Test3(t *testing.T) {
	input := "IX"
	except := 9
	actaul := romanToInt(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

// Explanation: L = 50, V = 5, III = 3.
func Test4(t *testing.T) {
	input := "LVIII"
	except := 58
	actaul := romanToInt(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
func Test5(t *testing.T) {
	input := "MCMXCIV"
	except := 1994
	actaul := romanToInt(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

func Test6(t *testing.T) {
	input := "IC"
	except := 99
	actaul := romanToInt(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

func Test7(t *testing.T) {
	input := "DCXXI"
	except := 621
	actaul := romanToInt(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}
