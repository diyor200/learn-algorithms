package divideANDconquer

import "testing"

func SumArrayWithRecursion(arr []int) int {
	sum := 0
	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		return arr[0]
	}

	return sum + SumArrayWithRecursion(arr[1:])
}

func Test_SumArrayWithRecursion(t *testing.T) {
	var arr = []int{2, 4, 6}
	result := SumArrayWithRecursion(arr)
	if result != 6 {
		t.Errorf("SumArrayWithRecursion expect 6, got %d", result)
		return
	}
	t.Log("ok")
}

func BinarySearchAlternativeUsingRecursion(arr []int, target int) bool {
	if len(arr) == 0 {
		return false
	} else if len(arr) == 1 {
		if arr[0] == target {
			return true
		} else {
			return false
		}
	}
	low, high := 0, len(arr)-1
	mid := (low + high) / 2
	if arr[mid] == target {
		return true
	}
	if arr[mid] > target {
		return BinarySearchAlternativeUsingRecursion(arr[:mid], target)
	}
	return BinarySearchAlternativeUsingRecursion(arr[mid+1:], target)
}

func Test_BinarySearchAlternativeUsingRecursion(t *testing.T) {
	var arr = []int{-10, -9, -1, 2, 4, 7}
	result := BinarySearchAlternativeUsingRecursion(arr, -9)
	if !result {
		t.Fatal("BinarySearchAlternativeUsingRecursion expect true, got false")
	}
	t.Log("ok")
}
