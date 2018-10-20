// Package longestsubstringwithoutrepeatingcharacters https://leetcode.com/problems/two-sum/description/
package longestsubstringwithoutrepeatingcharacters

/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring(s string) int {
	// Sample Solution, 4ms
	ans, i := 0, 0
	var arr [128]int
	for j, v := range s {
		i = max(arr[v], i)
		ans = max(ans, j-i+1)
		arr[v] = j + 1
	}
	return ans

	/* My Solution, 350ms
	result := 0
	for i, length := 0, len(s); i < length; i++ {
		wordMap := make(map[byte]bool)
		for j := i; j < length; j++ {
			if _, ok := wordMap[s[j]]; ok {
				break
			} else {
				wordMap[s[j]] = true
			}
		}
		cnt := len(wordMap)
		if result < cnt {
			result = cnt
		}
	}
	return result
	*/
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
