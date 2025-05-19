package main

import (
	"fmt"
)

// url: https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer/description/?envType=problem-list-v2&envId=binary-search

func main() {
	//var x = 7
	fmt.Println()
}

func maximumCount(nums []int) int {
	var zero, border int
	low, high := 0, len(nums)-1
	for {
		mid := (low + high) / 2
		if nums[mid] > 0 {
			if nums[mid-1] > 0 {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] == 0 {
			zero = mid
			break
		} else {
			high = mid - 1
		}
	}
}
