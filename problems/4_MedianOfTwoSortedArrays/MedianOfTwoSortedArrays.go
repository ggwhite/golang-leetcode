// Package medianofwwosortedarrays https://leetcode.com/problems/median-of-two-sorted-arrays/description/
package medianofwwosortedarrays

/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0

Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5

*/

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// merge two arrays
	nums := append(nums1, nums2...)

	// sort it
	sort.Ints(nums)

	length := len(nums)
	// when len % 2 == 0, median has two values, median value is average between these
	if length%2 == 0 {
		idx := length / 2
		return (float64(nums[idx-1]) + float64(nums[idx])) / 2
	}
	// when len % 2 != 0, median is single value
	// [1, 2, 3], len = 3, 3 / 2 = 1.5 (int is 1), median value index is 1
	idx := length / 2
	return float64(nums[idx])
}
