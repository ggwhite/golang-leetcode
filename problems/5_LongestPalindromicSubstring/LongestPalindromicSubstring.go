// Package longestpalindromicsubstring https://leetcode.com/problems/median-of-two-sorted-arrays/description/
package longestpalindromicsubstring

/*
https://leetcode.com/problems/longest-palindromic-substring/solution/

Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example 2:

Input: "cbbd"
Output: "bb"
*/

func longestPalindrome(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		start := max(i, len(ans))
		for j := start + 1; j <= len(s); j++ {
			p := s[i:j]
			if isPalindroe(p) && len(p) > len(ans) {
				ans = p
			}
		}
	}

	return ans
}

func isPalindroe(s string) bool {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return s == string(r)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
