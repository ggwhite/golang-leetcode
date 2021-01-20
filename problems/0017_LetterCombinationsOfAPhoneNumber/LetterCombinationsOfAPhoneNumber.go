// Package lettercombinationsofaphonenumber https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
package lettercombinationsofaphonenumber

/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.


Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.
*/

var mapping = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	// find total possible count
	// Ex: "237" => ["a", "b", "c"], ["d", "e", "f"], ["p", "q", "r", "s"], totla possible count is 3 * 3 * 4 = 36
	total := 1
	lengths := make([]int, len(digits))

	for i, r := range digits {
		// check input is 2-9
		if arr, ok := mapping[string(r)]; ok {
			lengths[i] = len(arr)
			total *= lengths[i]
		} else {
			return []string{}
		}
	}

	nodes := make([]string, total)

	for i := 0; i < total; i++ {
		var str string
		for j := 0; j < len(digits); j++ {
			key := string(digits[j])
			qty := 1
			for k := j + 1; k < len(digits); k++ {
				qty *= lengths[k]
			}
			str += mapping[key][(i/qty)%lengths[j]]
		}
		nodes[i] = str
	}
	return nodes
}

/**
Input "237"
Words Arr1 = ["a", "b", "c"], Arr2 = ["d", "e", "f"], Arr3 = ["p", "q", "r", "s"]
Length L1 = 3, L2 = 3, L3 = 4
Total possible count is 3 * 3 * 4 = 36

all possible is :

[0, 0, 0] (adp)	[1, 0, 0] (bdp) [2, 0, 0] (cdp)
[0, 0, 1] (adq)	[1, 0, 1] (bdq)	[2, 0, 1] (cdq)
[0, 0, 2] (adr)	[1, 0, 2] (bdr)	[2, 0, 2] (cdr)
[0, 0, 3] (ads)	[1, 0, 3] (bds)	[2, 0, 3] (cds)

[0, 1, 0] (aep)	[1, 1, 0] (bep) [2, 1, 0] (cep)
[0, 1, 1] (aeq)	[1, 1, 1] (beq)	[2, 1, 1] (ceq)
[0, 1, 2] (aer)	[1, 1, 2] (ber)	[2, 1, 2] (cer)
[0, 1, 3] (aes)	[1, 1, 3] (bes)	[2, 1, 3] (ces)

[0, 2, 0] (afp)	[1, 2, 0] (bfp) [2, 2, 0] (cfp)
[0, 2, 1] (afq)	[1, 2, 1] (bfq)	[2, 2, 1] (cfq)
[0, 2, 2] (afr)	[1, 2, 2] (bfr)	[2, 2, 2] (cfr)
[0, 2, 3] (afs)	[1, 2, 3] (bfs)	[2, 2, 3] (cfs)

for i = 0; i < total ; i++ {
	Word1 Index : (i / 1 * (L2 * L3)) % L1
	Word2 Index : (i / 1 * L3) % L2
	Word3 Index : (i / 1) % L3
}

i = 0
Word1 Index : (0 / 1 * (3 * 4)) % 3 = 0
Word2 Index : (0 / 1 * 4) % 3 = 0
Word3 Index : (0 / 1) % 4 = 0

i = 1
Word1 Index : (1 / 1 * (3 * 4)) % 3 = 0
Word2 Index : (1 / 1 * 4) % 3 = 0
Word3 Index : (1 / 1) % 4 = 1

i = 2
Word1 Index : (2 / 1 * (3 * 4)) % 3 = 0
Word2 Index : (2 / 1 * 4) % 3 = 0
Word3 Index : (2 / 1) % 4 = 2

i = 3
Word1 Index : (3 / 1 * (3 * 4)) % 3 = 0
Word2 Index : (3 / 1 * 4) % 3 = 0
Word3 Index : (3 / 1) % 4 = 3

i = 4
Word1 Index : (4 / 1 * (3 * 4)) % 3 = 0
Word2 Index : (4 / 1 * 4) % 3 = 1
Word3 Index : (4 / 1) % 4 = 0

...

*/
