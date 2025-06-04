package quicksort

import "testing"

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	var less, greater []int
	pivot := arr[0]

	for i := range arr {
		if arr[i] <= pivot {
			less = append(less, arr[i])
		} else {
			greater = append(greater, arr[i])
		}
	}

	return append(append(quicksort(less), pivot), quicksort(greater)...)
}

func TestQuickSort(t *testing.T) {
	arr := []int{10, 5, 2, 3, 9, 8, 1, 7}
	sorted := quicksort(arr)
	t.Log(sorted)
}
