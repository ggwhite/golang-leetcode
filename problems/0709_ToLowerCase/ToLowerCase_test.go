package tolowercase

import (
	"testing"
)

func Test1(t *testing.T) {
	input := "Hello"
	expect := "hello"
	actual := toLowerCase(input)
	if expect != actual {
		t.Errorf("Input: %s, Output: %s, But: %s", input, expect, actual)
	}
}

func Test2(t *testing.T) {
	input := "hello"
	expect := "hello"
	actual := toLowerCase(input)
	if expect != actual {
		t.Errorf("Input: %s, Output: %s, But: %s", input, expect, actual)
	}
}

func Test3(t *testing.T) {
	input := "LOVELY"
	expect := "lovely"
	actual := toLowerCase(input)
	if expect != actual {
		t.Errorf("Input: %s, Output: %s, But: %s", input, expect, actual)
	}
}
