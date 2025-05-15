package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	find := 10
	fmt.Println(BinarySearch(arr, find))
}

func BinarySearch(array []int, value int) int {
	var (
		low  int
		high = len(array) - 1
	)

	for low <= high {
		mid := (low + high) / 2
		if array[mid] == value {
			return mid
		} else if array[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

// EXERCISES
// 1.1 Suppose you have a sorted list of 128 names, and you’re searching
// through it using binary search. What’s the maximum number of
// steps it would take?
// Answer: 7

// 1.2 Suppose you double the size of the list. What’s the maximum
// number of steps now?
// Answer: increases by one step
