package uniquemorsecodewords

import (
	"testing"
)

func Test1(t *testing.T) {
	input := []string{"gin", "zen", "gig", "msg"}
	expect := 2
	actual := uniqueMorseRepresentations(input)
	if expect != actual {
		t.Errorf("Input: %v, Output: %d, But: %d", input, expect, actual)
	}
}
