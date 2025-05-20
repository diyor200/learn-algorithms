package main

import (
	"fmt"
)

// url: https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer/description/?envType=problem-list-v2&envId=binary-search

func main() {
	//var x = 7
	//var l = []int{-2, -1, -1, 1, 2, 3}
	//var l1 = []int{5, 20, 66, 1314}
	var l2 = []int{-3, -2, -1, 0, 0, 1, 2}
	//fmt.Println(maximumCount(l))
	//fmt.Println(maximumCount(l1))
	fmt.Println(maximumCount(l2))
}

func maximumCount(nums []int) int {
	if nums[0] > 0 {
		return len(nums)
	} else if nums[len(nums)-1] < 0 {
		return len(nums)
	}

	var negBorder, posBorder int
	low, high := 0, len(nums)-1
	for high >= low {
		mid := (low + high) / 2
		if nums[mid] > 0 {
			if nums[mid-1] > 0 {
				high = mid - 1
			} else {
				low = mid + 1
			}
			posBorder = mid
		} else if nums[mid] == 0 {
			if
			low = mid + 1
			zeroCount++
		} else {
			low = mid + 1
		}
	}

	return max(len(nums[:posBorder-1])-zeroCount, len(nums[posBorder:]))
}
