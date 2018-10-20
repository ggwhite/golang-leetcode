package longestsubstringwithoutrepeatingcharacters

import (
	"testing"
)

func Test1(t *testing.T) {
	input := "abcabcbb"
	expect := 3
	actual := lengthOfLongestSubstring(input)
	if expect != actual {
		t.Errorf("Input: %s, Output: %d, But: %d", input, expect, actual)
	}
}

func Test2(t *testing.T) {
	input := "bbbbb"
	expect := 1
	actual := lengthOfLongestSubstring(input)
	if expect != actual {
		t.Errorf("Input: %s, Output: %d, But: %d", input, expect, actual)
	}
}

func Test3(t *testing.T) {
	input := "pwwkew"
	expect := 3
	actual := lengthOfLongestSubstring(input)
	if expect != actual {
		t.Errorf("Input: %s, Output: %d, But: %d", input, expect, actual)
	}
}

func Test4(t *testing.T) {
	input := " "
	expect := 1
	actual := lengthOfLongestSubstring(input)
	if expect != actual {
		t.Errorf("Input: %s, Output: %d, But: %d", input, expect, actual)
	}
}
