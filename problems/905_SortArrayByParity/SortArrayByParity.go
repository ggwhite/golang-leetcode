// Package sortarraybyparity https://leetcode.com/problems/sort-array-by-parity/description/
package sortarraybyparity

/*
Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.

You may return any answer array that satisfies this condition.

Example 1:

Input: [3,1,2,4]
Output: [2,4,3,1]
The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

Note:

1 <= A.length <= 5000
0 <= A[i] <= 5000
*/

func sortArrayByParity(A []int) []int {
	/*
		// two slice of 『even』 and 『odd』, return 『even』 + 『odd』, 70 ms
		even := []int{}
		odd := []int{}
		for _, el := range A {
			if el%2 == 0 {
				even = append(even, el)
			} else {
				odd = append(odd, el)
			}
		}
		return append(even, odd...)
	*/

	// prepend 『even』, append 『odd』, 100 ms
	result := []int{}

	for _, el := range A {
		if el%2 == 0 {
			result = append([]int{el}, result...)
		} else {
			result = append(result, el)
		}
	}

	return result
}
