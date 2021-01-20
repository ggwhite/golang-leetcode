package lettercombinationsofaphonenumber

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	input := "23"
	except := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	actaul := letterCombinations(input)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %s, Output:%v, But: %v", input, except, actaul)
	}
}

func Test2(t *testing.T) {
	input := "237"
	except := []string{"adp", "adq", "adr", "ads", "aep", "aeq", "aer", "aes", "afp", "afq", "afr", "afs", "bdp", "bdq", "bdr", "bds", "bep", "beq", "ber", "bes", "bfp", "bfq", "bfr", "bfs", "cdp", "cdq", "cdr", "cds", "cep", "ceq", "cer", "ces", "cfp", "cfq", "cfr", "cfs"}
	actaul := letterCombinations(input)
	if !reflect.DeepEqual(except, actaul) {
		t.Errorf("Input: %s, Output:%v, But: %v", input, except, actaul)
	}
}
