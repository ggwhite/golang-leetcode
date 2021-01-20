package zigzagconversion

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	input := "PAYPALISHIRING"
	numRows := 3
	expect := "PAHNAPLSIIGYIR"
	actual := convert(input, numRows)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("Input: %s, NumRows: %d, Output: %s, But: %s", input, numRows, expect, actual)
	}
}

func Test2(t *testing.T) {
	input := "PAYPALISHIRING"
	numRows := 4
	expect := "PINALSIGYAHRPI"
	actual := convert(input, numRows)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("Input: %s, NumRows: %d, Output: %s, But: %s", input, numRows, expect, actual)
	}
}
