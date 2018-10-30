// Package palindromenumber https://leetcode.com/problems/palindrome-number/description/
package palindromenumber

/*
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true

Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

Follow up:

Coud you solve it without converting the integer to a string?
*/

func isPalindrome(x int) bool {
	// Sample 56ms
	if x < 0 {
		return false
	}
	y := x
	ans := 0
	for x > 0 {
		ans = ans*10 + (x % 10)
		x = x / 10
	}
	return ans == y

	/* My Solution 60ms
	if x < 0 {
		return false
	}
	deg := int(math.Log10(float64(x))) + 1
	if deg == 1 {
		return true
	}
	if deg <= 2 {
		return int(x/10) == int(x%10)
	}
	var dl, dr int
	if deg%2 == 0 {
		dl = deg / 2
		dr = dl - 1
		median := (x/int(math.Pow10(dl)))%10*10 + (x/int(math.Pow10(dr)))%10
		if !isPalindrome(median) {
			return false
		}
	} else {
		dl = deg / 2
		dr = deg / 2
	}
	for i, j := dl, dr; i < deg; i, j = i+1, j-1 {
		left := x / int(math.Pow10(i)) % 10
		right := x / int(math.Pow10(j)) % 10
		if left != right {
			return false
		}
	}
	return true
	*/
}
