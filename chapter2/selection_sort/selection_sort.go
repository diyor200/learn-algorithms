package main

import "fmt"

func main() {
	arr := []int{64, 25, 12, 12, 22, 11}
	fmt.Println("Before sorting:", arr)

	selectionSort(arr)

	fmt.Println("After sorting:", arr)
}

func selectionSort(nums []int) {
	n := len(nums)

	for i := 0; i < n-1; i++ {
		minIndex := i

		for j := i + 1; j < n; j++ {
			if nums[minIndex] > nums[j] {
				minIndex = j
			}
		}

		if minIndex != i {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}
}
