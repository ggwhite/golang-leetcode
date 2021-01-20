// Package validparentheses https://leetcode.com/problems/valid-parentheses/description/
package validparentheses

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

	1. Open brackets must be closed by the same type of brackets.
	2. Open brackets must be closed in the correct order.

Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true

Example 2:

Input: "()[]{}"
Output: true

Example 3:

Input: "(]"
Output: false

Example 4:

Input: "([)]"
Output: false

Example 5:

Input: "{[]}"
Output: true
*/

func isValid(s string) bool {
	// Sample 0ms, Use Stack
	var stack []byte
	endmap := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			continue
		}
		sn := len(stack)
		if sn == 0 {
			return false
		}
		if endmap[stack[sn-1]] != s[i] {
			return false
		}
		stack = stack[:sn-1]
	}
	return len(stack) == 0

	/* My Solution 16ms
	n := len(s)
	if n == 0 {
		return true
	}
	if n == 1 {
		return false
	}
	if n == 2 && s != "()" && s != "[]" && s != "{}" {
		return false
	}
	// is contain others return false
	for i := 0; i < n; i++ {
		if s[i] != '(' && s[i] != ')' && s[i] != '[' && s[i] != ']' && s[i] != '{' && s[i] != '}' {
			return false
		}
	}

	var newstr string
	newstr = s
	newstr = strings.Replace(newstr, "()", "", -1)
	newstr = strings.Replace(newstr, "[]", "", -1)
	newstr = strings.Replace(newstr, "{}", "", -1)

	// check if nothing replaced
	if s == newstr {
		return false
	}
	return isValid(newstr)
	*/
}
