package jewelsandstones

import (
	"testing"
)

func Test1(t *testing.T) {
	J := "aA"
	S := "aAAbbbb"
	expect := 3
	actual := numJewelsInStones(J, S)
	if expect != actual {
		t.Errorf("Input: J=%s, S=%s, Output: %d, But: %d", J, S, expect, actual)
	}
}

func Test2(t *testing.T) {
	J := "z"
	S := "ZZ"
	expect := 0
	actual := numJewelsInStones(J, S)
	if expect != actual {
		t.Errorf("Input: J=%s, S=%s, Output: %d, But: %d", J, S, expect, actual)
	}
}
