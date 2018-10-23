package stringtointeger

import (
	"testing"
)

func Test1(t *testing.T) {
	input := "42"
	except := 42
	actaul := myAtoi(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

// Explanation:
// The first non-whitespace character is '-', which is the minus sign.
// Then take as many numerical digits as possible, which gets 42.
func Test2(t *testing.T) {
	input := "   -42"
	except := -42
	actaul := myAtoi(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

// Explanation:
// Conversion stops at digit '3' as the next character is not a numerical digit.
func Test3(t *testing.T) {
	input := "4193 with words"
	except := 4193
	actaul := myAtoi(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

// Explanation:
// The first non-whitespace character is 'w', which is not a numerical
// digit or a +/- sign. Therefore no valid conversion could be performed.
func Test4(t *testing.T) {
	input := "words and 987"
	except := 0
	actaul := myAtoi(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

// Explanation:
// The number "-91283472332" is out of the range of a 32-bit signed integer.
// Thefore INT_MIN (−231) is returned.
func Test5(t *testing.T) {
	input := "-91283472332"
	except := -2147483648
	actaul := myAtoi(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

// Explanation:
// The number "32311143423" is out of the range of a 32-bit signed integer.
// Thefore INT_MIN (231 - 1) is returned.
func Test6(t *testing.T) {
	input := "32311143423"
	except := 2147483647
	actaul := myAtoi(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

func Test7(t *testing.T) {
	input := "          "
	except := 0
	actaul := myAtoi(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

func Test8(t *testing.T) {
	input := "+42"
	except := 42
	actaul := myAtoi(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

func Test9(t *testing.T) {
	input := ".2"
	except := 0
	actaul := myAtoi(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

func Test10(t *testing.T) {
	input := ""
	except := 0
	actaul := myAtoi(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

func Test11(t *testing.T) {
	input := "  -0012a42"
	except := -12
	actaul := myAtoi(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

func Test12(t *testing.T) {
	input := "3231114342x3"
	except := 2147483647
	actaul := myAtoi(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

func Test13(t *testing.T) {
	input := "-9128347233a2"
	except := -2147483648
	actaul := myAtoi(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

func Test14(t *testing.T) {
	input := "9223372036854775808"
	except := 2147483647
	actaul := myAtoi(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}

func Test15(t *testing.T) {
	input := "-9223372036854775808"
	except := -2147483648
	actaul := myAtoi(input)
	if except != actaul {
		t.Errorf("Input: %s, Output:%d, But: %d", input, except, actaul)
	}
}
