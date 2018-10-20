// Package twosum https://leetcode.com/problems/two-sum/description/
package twosum

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

func twoSum(nums []int, target int) []int {
	// Sample
	indexMap := make(map[int]int)

	for index, num := range nums {
		rest := target - num
		index2, ok := indexMap[rest]
		if ok && index2 != index {
			return []int{index2, index}
		}
		indexMap[num] = index
	}

	/* My solution
	for i, len := 0, len(nums); i < len; i++ {
		for j := i + 1; j < len; j++ {
			if sum := nums[i] + nums[j]; sum == target {
				return []int{i, j}
			}
		}
	}
	*/
	return nil
}
