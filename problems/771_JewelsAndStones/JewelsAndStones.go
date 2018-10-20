// Package jewelsandstones https://leetcode.com/problems/jewels-and-stones/description/
package jewelsandstones

/*
You're given strings J representing the types of stones that are jewels, and S representing the stones you have.
Each character in S is a type of stone you have.
You want to know how many of the stones you have are also jewels.

The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".

Example 1:

Input: J = "aA", S = "aAAbbbb"
Output: 3

Example 2:

Input: J = "z", S = "ZZ"
Output: 0

Note:

S and J will consist of letters and have length at most 50.
The characters in J are distinct.
*/

func numJewelsInStones(J string, S string) int {
	sum := 0

	charMap := make(map[string]int)

	for i, len := 0, len(S); i < len; i++ {
		charMap[S[i:i+1]]++
	}

	for i, len := 0, len(J); i < len; i++ {
		if value, ok := charMap[J[i:i+1]]; ok {
			sum += value
		}
	}

	return sum
}
