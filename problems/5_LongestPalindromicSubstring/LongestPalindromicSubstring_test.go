package longestpalindromicsubstring

import (
	"reflect"
	"testing"
)

// func Test1(t *testing.T) {
// 	input := "babad"
// 	expect := "bab"
// 	actual := longestPalindrome(input)

// 	if !reflect.DeepEqual(expect, actual) {
// 		t.Errorf("Input: %s, Output: %s, But: %s", input, expect, actual)
// 	}
// }

// func Test2(t *testing.T) {
// 	input := "cbbd"
// 	expect := "bb"
// 	actual := longestPalindrome(input)

// 	if !reflect.DeepEqual(expect, actual) {
// 		t.Errorf("Input: %s, Output: %s, But: %s", input, expect, actual)
// 	}
// }

// func Test3(t *testing.T) {
// 	input := "cbddbcd"
// 	expect := "cbddbc"
// 	actual := longestPalindrome(input)

// 	if !reflect.DeepEqual(expect, actual) {
// 		t.Errorf("Input: %s, Output: %s, But: %s", input, expect, actual)
// 	}
// }

func Test3(t *testing.T) {
	input := "babadada"
	expect := "adada"
	actual := longestPalindrome(input)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("Input: %s, Output: %s, But: %s", input, expect, actual)
	}
}

func Test4(t *testing.T) {
	input := "zudfweormatjycujjirzjpyrmaxurectxrtqedmmgergwdvjmjtstdhcihacqnothgttgqfywcpgnuvwglvfiuxteopoyizgehkwuvvkqxbnufkcbodlhdmbqyghkojrgokpwdhtdrwmvdegwycecrgjvuexlguayzcammupgeskrvpthrmwqaqsdcgycdupykppiyhwzwcplivjnnvwhqkkxildtyjltklcokcrgqnnwzzeuqioyahqpuskkpbxhvzvqyhlegmoviogzwuiqahiouhnecjwysmtarjjdjqdrkljawzasriouuiqkcwwqsxifbndjmyprdozhwaoibpqrthpcjphgsfbeqrqqoqiqqdicvybzxhklehzzapbvcyleljawowluqgxxwlrymzojshlwkmzwpixgfjljkmwdtjeabgyrpbqyyykmoaqdambpkyyvukalbrzoyoufjqeftniddsfqnilxlplselqatdgjziphvrbokofvuerpsvqmzakbyzxtxvyanvjpfyvyiivqusfrsufjanmfibgrkwtiuoykiavpbqeyfsuteuxxjiyxvlvgmehycdvxdorpepmsinvmyzeqeiikajopqedyopirmhymozernxzaueljjrhcsofwyddkpnvcvzixdjknikyhzmstvbducjcoyoeoaqruuewclzqqqxzpgykrkygxnmlsrjudoaejxkipkgmcoqtxhelvsizgdwdyjwuumazxfstoaxeqqxoqezakdqjwpkrbldpcbbxexquqrznavcrprnydufsidakvrpuzgfisdxreldbqfizngtrilnbqboxwmwienlkmmiuifrvytukcqcpeqdwwucymgvyrektsnfijdcdoawbcwkkjkqwzffnuqituihjaklvthulmcjrhqcyzvekzqlxgddjoir"
	expect := "gykrkyg"
	actual := longestPalindrome(input)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("Input: %s, Output: %s, But: %s", input, expect, actual)
	}
}
