// Package generateparentheses https://leetcode.com/problems/generate-parentheses/description/
package generateparentheses

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

func generateParenthesis(n int) []string {
	var ans []string
	ans = backtrack(ans, "", 0, 0, n)
	return ans
}

func backtrack(ans []string, cur string, open int, close int, max int) []string {
	// point a
	if len(cur) == max*2 {
		ans = append(ans, cur)
		return ans
	}
	// point b
	if open < max {
		ans = backtrack(ans, cur+"(", open+1, close, max)
	}
	// point c
	if close < open {
		ans = backtrack(ans, cur+")", open, close+1, max)
	}
	return ans
}

/*
n = 3, init ans = []
ans, "", 0, 0, 3
	a	0 == 3*2 -> false
	b	0 < 3 -> true
		ans, "(", 1, 0, 3
			a	1 == 3*2 -> false
			b	1 < 3 -> true
				ans, "((", 2, 0, 3
					a	2 == 3*2 -> false
					b	2 < 3 -> true
						ans, "(((", 3, 0, 3
							a	3 != 3*2 -> false
							b	3 < 3 -> false
							c	0 < 3 -> true
								ans, "((()", 3, 1, 3
									a	4 == 3*2 -> false
									b	3 < 3 -> false
									c	1 < 3 -> true
										ans, "((())", 3, 2, 3
											a	5 == 3*2 -> false
											b	3 < 3 -> false
											c	2 < 3 -> true
												ans, "((()))", 3, 3, 3
													a	6 == 3*2 -> true
														ans append ((())) return
					c	0 < 2 -> true
						ans, "(()", 2, 1, 3
							a	3 != 3*2 -> false
							b	2 < 3 -> true
								ans, "(()(", 3, 1, 3
									a	4 != 3*2 -> false
									b	3 < 3 -> false
									c	1 < 3 -> true
										ans, "(()()", 3, 2, 3
											a	5 != 3*2 -> false
											b	3 < 3 -> false
											c	2 < 3 -> true
												ans, "(()())", 3, 3, 3
													a	6 == 3*2 -> true
														ans append (()()) return
							c 	1 < 2 -> true
								ans, "(())", 2, 2, 3
									a	4 != 3*2 -> false
									b	ans, "(())(", 3, 2, 3
											a	5 != 3*2 -> false
											b	3 < 3 -> false
											c	2 < 3 -> true
												ans, "(())()", 3, 3, 3
													a	6 == 3*2 -> true
														ans append (())() return
									c	2 < 2 -> false
			c	0 < 1 -> true
				ans, "()", 1, 1, 3
					a	2 == 3*2 -> false
					b	1 < 3 -> true
						ans, "()(", 2, 1, 3
							a	3 == 3*2 -> false
							b	2 < 3 -> true
								ans, "()((", 3, 1, 3
									a	4 == 3*2 -> false
									b	3 < 3 -> false
									c	1 < 2 -> true
										ans, "()(()", 3, 2, 3
											a	5 == 3*2 -> false
											b	3 < 3 -> false
											c	2 < 3 -> true
												ans, "()(())", 3, 3, 3
													a	6 == 3*2 -> true
														ans append ()(()) return
							c	1 < 2 -> true
								ans, "()()", 2, 2, 3
									a	4 == 3*2 -> false
									b	2 < 3 -> true
										ans, "()()(", 3, 2, 3
											a	5 == 3*2 -> false
											b	3 < 3 -> false
											c	2 < 3 -> true
												ans, "()()()", 3, 3, 3
													a	6 == 3*2 -> true
														ans append ()()() return
									c	2 < 2 -> false
					c	1 < 1 -> false
	c	0 < 0 -> false
return ans
*/
