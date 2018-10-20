// Package tolowercase https://leetcode.com/problems/to-lower-case/description/
package tolowercase

/*
Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.

Example 1:

Input: "Hello"
Output: "hello"

Example 2:

Input: "here"
Output: "here"

Example 3:

Input: "LOVELY"
Output: "lovely"
*/

import "strings"

func toLowerCase(str string) string {
	return strings.ToLower(str)
}
