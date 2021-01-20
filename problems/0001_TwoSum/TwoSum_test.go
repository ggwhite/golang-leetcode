package twosum

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	expect := []int{0, 1}
	actual := twoSum(nums, 9)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("nums: %v, expect: %v, actual: %v", nums, expect, actual)
	}
}
