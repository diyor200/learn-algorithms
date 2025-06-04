package leetcode

import (
	"fmt"
	"testing"
)

// url: https://leetcode.com/problems/majority-element/?envType=problem-list-v2&envId=sorting
func majorityElement(nums []int) int {
	quicksort(nums, 0, len(nums)-1)
	return nums[len(nums)/2]
}

func quicksort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quicksort(arr, low, p-1)
		quicksort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func TestMajorityElement(t *testing.T) {
	//var nums = []int{2, 2, 1, 1, 1, 2, 2}
	//var nums = []int{2, 1, 2}
	//var nums = []int{2, 2}
	var nums = []int{1}
	fmt.Println(majorityElement(nums)) //
}

// url: https://leetcode.com/problems/sort-an-array/
func sortArray(nums []int) []int {
	quicksort(nums, 0, len(nums)-1)
	return nums
}

func TestSortArray(t *testing.T) {
	var nums = []int{5, 2, 3, 1}
	fmt.Println(sortArray(nums))
}
