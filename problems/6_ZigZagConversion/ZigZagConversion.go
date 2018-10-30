// Package zigzagconversion https://leetcode.com/problems/zigzag-conversion/description/
package zigzagconversion

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);

Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"

Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
*/

func convert(s string, numRows int) string {
	// Sample 8ms
	if numRows < 2 {
		return s
	}
	n := len(s)
	out := make([]byte, n)
	idx := 0
	for i := 0; i < numRows; i++ {
		zig := (numRows - i - 1) * 2
		zag := i * 2
		if zig == 0 {
			zig = zag
		}
		if zag == 0 {
			zag = zig
		}
		for j, k := i, 0; j < n; idx, k = idx+1, k+1 {
			out[idx] = s[j]
			if k%2 == 0 {
				j += zig
			} else {
				j += zag
			}
		}
	}

	return string(out)

	/* My Solution 24ms
	if numRows < 2 {
		return s
	}
	var ans []rune
	r := make([][]rune, numRows)
	plen := numRows*2 - 2
	for i, c := range s {
		idx := i % plen
		if idx >= numRows {
			idx = plen - idx
		}
		r[idx] = append(r[idx], c)
	}
	for _, row := range r {
		ans = append(ans, row...)
	}
	return string(ans)
	*/
}
