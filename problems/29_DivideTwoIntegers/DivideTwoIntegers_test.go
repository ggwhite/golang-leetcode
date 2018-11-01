package dividetwointegers

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	dividend := 10
	divisor := 3
	except := 3
	actaul := divide(dividend, divisor)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("haystack: %d, divisor: %d, Output: %d, But: %d", dividend, divisor, except, actaul)
	}
}

func Test2(t *testing.T) {
	dividend := 7
	divisor := -3
	except := -2
	actaul := divide(dividend, divisor)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("dividend: %d, divisor: %d, Output: %d, But: %d", dividend, divisor, except, actaul)
	}
}
