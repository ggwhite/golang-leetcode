package implementstr

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	haystack := "hello"
	needle := "ll"
	except := 2
	actaul := strStr(haystack, needle)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("haystack: %s, needle: %s, Output: %d, But: %d", haystack, needle, except, actaul)
	}
}

func Test2(t *testing.T) {
	haystack := "aaaaa"
	needle := "bba"
	except := -1
	actaul := strStr(haystack, needle)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("haystack: %s, needle: %s, Output: %d, But: %d", haystack, needle, except, actaul)
	}
}
