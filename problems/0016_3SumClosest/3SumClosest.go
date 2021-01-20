// Package threesumclosest https://leetcode.com/problems/3sum-closest/description/
package threesumclosest

/*
Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target.
Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:

Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
*/

import "sort"

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	sort.Ints(nums)

	// check target overed limit
	max := nums[n-1] + nums[n-2] + nums[n-3]
	min := nums[0] + nums[1] + nums[2]

	if target >= max {
		return max
	}
	if target <= min {
		return min
	}

	var ans int

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return sum
			}
			// check sum & ans is closest
			if (target-sum)*(target-sum) <= (target-ans)*(target-ans) {
				ans = sum
			}
			if sum > target {
				// find smaller sum
				for r--; l < r && nums[r] == nums[r+1]; r-- {
				}
			} else {
				// find bigger sum
				for l++; l < r && nums[l] == nums[l-1]; l++ {
				}
			}
		}
	}
	return ans
}
