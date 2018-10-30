package integertoroman

import (
	"testing"
)

func Test1(t *testing.T) {
	input := 3
	except := "III"
	actaul := intToRoman(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%s, But: %s", input, except, actaul)
	}
}

func Test2(t *testing.T) {
	input := 4
	except := "IV"
	actaul := intToRoman(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%s, But: %s", input, except, actaul)
	}
}

func Test3(t *testing.T) {
	input := 9
	except := "IX"
	actaul := intToRoman(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%s, But: %s", input, except, actaul)
	}
}

// Explanation: L = 50, V = 5, III = 3.
func Test4(t *testing.T) {
	input := 58
	except := "LVIII"
	actaul := intToRoman(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%s, But: %s", input, except, actaul)
	}
}

// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
func Test5(t *testing.T) {
	input := 1994
	except := "MCMXCIV"
	actaul := intToRoman(input)
	if except != actaul {
		t.Errorf("Input: %d, Output:%s, But: %s", input, except, actaul)
	}
}
