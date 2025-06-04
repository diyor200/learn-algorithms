package leetcode

import "testing"

// url: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/?envType=problem-list-v2&envId=array
func maxProfit(prices []int) int {
	minPrice := int(^uint(0) >> 1)
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

func TestMaxProfit(t *testing.T) {
	var prices = []int{7, 1, 5, 3, 6, 4}
	var prices1 = []int{7, 6, 4, 3, 1}
	if maxProfit(prices) != 5 {
		t.Fail()
	}

	if maxProfit(prices1) != 0 {
		t.Fail()
	}

	t.Log("pass")
}

// Url: https://leetcode.com/problems/merge-sorted-array/?envType=problem-list-v2&envId=array
func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums2) == 0 {
		return
	}

	for i := 0; i < len(nums2); i++ {
		nums1[m+i] = nums2[i]
	}
	nums2 = nil

	for i := 0; i < m+n; i++ {
		for j := i + 1; j < m+n; j++ {
			if nums1[i] > nums1[j] {
				nums1[i], nums1[j] = nums1[j], nums1[i]
			}
		}
	}
}

func mergeGPT(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}

		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	t.Log(nums1)
}

// url: https://leetcode.com/problems/single-number/description/?envType=problem-list-v2&envId=array
func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var res int
	for i := range nums {
		res ^= nums[i]
	}

	//var m = make(map[int]int, len(nums)/2+1)
	//for i := range nums {
	//	m[nums[i]]++
	//}
	//
	//var res int
	//for i := range m {
	//	if m[i] == 1 {
	//		res = i
	//		break
	//	}
	//}
	//
	//m = nil

	return res
}

func TestSingleNumber(t *testing.T) {
	arr := []int{4, 1, 2, 1, 2}
	t.Log(singleNumber(arr))
}
