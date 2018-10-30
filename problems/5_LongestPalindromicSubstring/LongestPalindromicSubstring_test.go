package longestpalindromicsubstring

import (
	"reflect"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	input := "babad"
	expect := "aba"
	actual := longestPalindrome(input)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("Input: %s, Output: %s, But: %s", input, expect, actual)
	}
}

func Test2(t *testing.T) {
	input := "cbbd"
	expect := "bb"
	actual := longestPalindrome(input)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("Input: %s, Output: %s, But: %s", input, expect, actual)
	}
}

func Test3(t *testing.T) {
	input := "cbddbcd"
	expect := "cbddbc"
	actual := longestPalindrome(input)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("Input: %s, Output: %s, But: %s", input, expect, actual)
	}
}

func Test4(t *testing.T) {
	input := "babadada"
	expect := "adada"
	actual := longestPalindrome(input)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("Input: %s, Output: %s, But: %s", input, expect, actual)
	}
}

func Test5(t *testing.T) {
	input := "zudfweormatjycujjirzjpyrmaxurectxrtqedmmgergwdvjmjtstdhcihacqnothgttgqfywcpgnuvwglvfiuxteopoyizgehkwuvvkqxbnufkcbodlhdmbqyghkojrgokpwdhtdrwmvdegwycecrgjvuexlguayzcammupgeskrvpthrmwqaqsdcgycdupykppiyhwzwcplivjnnvwhqkkxildtyjltklcokcrgqnnwzzeuqioyahqpuskkpbxhvzvqyhlegmoviogzwuiqahiouhnecjwysmtarjjdjqdrkljawzasriouuiqkcwwqsxifbndjmyprdozhwaoibpqrthpcjphgsfbeqrqqoqiqqdicvybzxhklehzzapbvcyleljawowluqgxxwlrymzojshlwkmzwpixgfjljkmwdtjeabgyrpbqyyykmoaqdambpkyyvukalbrzoyoufjqeftniddsfqnilxlplselqatdgjziphvrbokofvuerpsvqmzakbyzxtxvyanvjpfyvyiivqusfrsufjanmfibgrkwtiuoykiavpbqeyfsuteuxxjiyxvlvgmehycdvxdorpepmsinvmyzeqeiikajopqedyopirmhymozernxzaueljjrhcsofwyddkpnvcvzixdjknikyhzmstvbducjcoyoeoaqruuewclzqqqxzpgykrkygxnmlsrjudoaejxkipkgmcoqtxhelvsizgdwdyjwuumazxfstoaxeqqxoqezakdqjwpkrbldpcbbxexquqrznavcrprnydufsidakvrpuzgfisdxreldbqfizngtrilnbqboxwmwienlkmmiuifrvytukcqcpeqdwwucymgvyrektsnfijdcdoawbcwkkjkqwzffnuqituihjaklvthulmcjrhqcyzvekzqlxgddjoir"
	expect := "gykrkyg"

	begin := time.Now()
	actual := longestPalindrome(input)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("Input: %s, Output: %s, But: %s", input, expect, actual)
	}

	if time.Since(begin) > 500*time.Millisecond {
		t.Errorf("Input: %s, Output: %s, Time Limit Exceeded", input, expect)
	}
}
