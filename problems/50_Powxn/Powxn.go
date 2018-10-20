// Package powxn https://leetcode.com/problems/powx-n/description/
package powxn

/*
Implement pow(x, n), which calculates x raised to the power n (xn).

Example 1:

Input: 2.00000, 10
Output: 1024.00000

Example 2:

Input: 2.10000, 3
Output: 9.26100

Example 3:

Input: 2.00000, -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25

Note:

-100.0 < x < 100.0
n is a 32-bit signed integer, within the range [−231, 231 − 1]
*/

import "math"

func myPow(x float64, n int) float64 {
	/* Sample 0ms
		if n == 0 {
			return 1.0
		}

		var unsign_n uint32
		var curPowX float64
	    var ans float64 = 1

		if n < 0 {
			curPowX = 1 / x
			unsign_n = uint32(-n)
		} else {
			curPowX = x
			unsign_n = uint32(n)
		}

		for unsign_n != 0 {
			if unsign_n&0x01 == 0x01 {
				ans *= curPowX
			}
			curPowX *= curPowX
			unsign_n >>= 1
		}
		return ans
	*/

	// My Solution 0ms
	return math.Pow(x, float64(n))
}
