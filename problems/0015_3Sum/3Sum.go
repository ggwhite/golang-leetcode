// Package threesum https://leetcode.com/problems/3sum/description/
package threesum

/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0?
Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

import "sort"

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	n := len(nums)
	if n < 3 {
		return ans
	}
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]}) // add these numbers to anser
				l, r = l+1, r-1                                     // find next set
				for ; l < r && nums[l] == nums[l-1]; l++ {
				} // skip the same number at this time from left
				for ; l < r && nums[r] == nums[r+1]; r-- {
				} // skip the same number at this time from right
			} else if sum < 0 { // left ++ can find bigger number
				l++
			} else { // right -- can find smaller number
				r--
			}
		}
	}
	return ans
}
