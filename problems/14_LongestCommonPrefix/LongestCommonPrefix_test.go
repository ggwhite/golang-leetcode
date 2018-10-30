package longestcommonprefix

import (
	"testing"
)

func Test1(t *testing.T) {
	input := []string{"flower", "flow", "flight"}
	except := "fl"
	actaul := longestCommonPrefix(input)
	if except != actaul {
		t.Errorf("Input: %v, Output:%s, But: %s", input, except, actaul)
	}
}

func Test2(t *testing.T) {
	input := []string{"dog", "racecar", "car"}
	except := ""
	actaul := longestCommonPrefix(input)
	if except != actaul {
		t.Errorf("Input: %v, Output:%s, But: %s", input, except, actaul)
	}
}

func Test3(t *testing.T) {
	input := []string{}
	except := ""
	actaul := longestCommonPrefix(input)
	if except != actaul {
		t.Errorf("Input: %v, Output:%s, But: %s", input, except, actaul)
	}
}

func Test4(t *testing.T) {
	input := []string{"a"}
	except := "a"
	actaul := longestCommonPrefix(input)
	if except != actaul {
		t.Errorf("Input: %v, Output:%s, But: %s", input, except, actaul)
	}
}

func Test5(t *testing.T) {
	input := []string{"abc"}
	except := "abc"
	actaul := longestCommonPrefix(input)
	if except != actaul {
		t.Errorf("Input: %v, Output:%s, But: %s", input, except, actaul)
	}
}

func Test6(t *testing.T) {
	input := []string{"c", "c"}
	except := "c"
	actaul := longestCommonPrefix(input)
	if except != actaul {
		t.Errorf("Input: %v, Output:%s, But: %s", input, except, actaul)
	}
}
