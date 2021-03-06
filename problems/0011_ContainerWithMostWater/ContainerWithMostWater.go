// Package containerwithmostwater https://leetcode.com/problems/container-with-most-water/description/
package containerwithmostwater

/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.





The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.



Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49
*/

func maxArea(height []int) int {
	var ans int
	for i, j := 0, len(height)-1; i < j; {
		w := j - i
		h := min(height[i], height[j])
		ans = max(ans, w*h)
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
