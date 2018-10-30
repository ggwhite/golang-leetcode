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
	/* My Org Solution
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
	*/

	// After Read Solution: Use Expand Around Center
	if len(s) == 0 {
		return s
	}
	var start, end int
	for i := 0; i < len(s); i++ {
		expand := max(expandAroundCenter(s, i, i), expandAroundCenter(s, i, i+1))
		if expand > end-start {
			start = i - (expand-1)/2
			end = i + expand/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
