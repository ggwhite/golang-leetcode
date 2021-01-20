package medianofwwosortedarrays

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	expect := float64(2)
	actual := findMedianSortedArrays(nums1, nums2)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("nums1: %v, nums2: %v, expect: %v, actual: %v", nums1, nums2, expect, actual)
	}
}

func Test2(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	expect := 2.5
	actual := findMedianSortedArrays(nums1, nums2)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("nums1: %v, nums2: %v, expect: %v, actual: %v", nums1, nums2, expect, actual)
	}
}
