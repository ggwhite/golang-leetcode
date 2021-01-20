// Package implementstr https://leetcode.com/problems/implement-strstr/description/
package implementstr

/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
*/

func strStr(haystack string, needle string) int {
	n1, n2 := len(haystack), len(needle)
	if n2 == 0 {
		return 0
	}
	for i := 0; i < n1-n2+1; i++ {
		if haystack[i:i+n2] == needle {
			return i
		}
	}
	return -1
}
