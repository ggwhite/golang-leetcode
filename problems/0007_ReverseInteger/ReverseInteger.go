// Package reverseinteger https://leetcode.com/problems/reverse-integer/description/
package reverseinteger

/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321

Example 2:

Input: -123
Output: -321

Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

import "math"

func reverse(x int) int {
	if overflow(x) {
		return 0
	}
	var ans int
	for math.Abs(float64(x)) >= 1 {
		ans = ans*10 + x%10
		x = x / 10
	}
	if overflow(ans) {
		return 0
	}
	return ans
}

func overflow(x int) bool {
	return x < -1<<31 || x > 1<<31-1
}
