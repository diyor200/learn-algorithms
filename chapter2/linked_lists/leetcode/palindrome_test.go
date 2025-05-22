package leetcode

import (
	"testing"
)

// url: https://leetcode.com/problems/palindrome-linked-list/description/?envType=problem-list-v2&envId=linked-list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	l := lenList(head)
	var h = head
	for i := 0; i < l; i++ {
		if head.Val != getNode(l-1-i, h).Val {
			return false
		}

		head = head.Next
	}

	return true
}

func lenList(head *ListNode) int {
	var res int
	for {
		res++
		if head.Next == nil {
			break
		} else {
			head = head.Next
		}
	}

	return res
}

func getNode(index int, head *ListNode) *ListNode {
	for i := 1; i < index; i++ {
		head = head.Next
	}

	return head
}

func reverse(head *ListNode) *ListNode {
	reversed := &ListNode{
		Val: head.Val,
	}

	for {
		if head.Next == nil {
			return reversed
		}

		head = head.Next

		reversed = &ListNode{
			Val:  head.Val,
			Next: reversed,
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	var head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}

	if isPalindrome(head) == true {
		t.Log("pass")
	} else {
		t.Fail()
	}
}

// URL: https://leetcode.com/problems/remove-nth-node-from-end-of-list/?envType=problem-list-v2&envId=linked-list
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := lenList(head)
	if length == 1 || length == 0 {
		return nil
	} else if length == n {
		return head.Next
	}

	nThNode := getNode(length-(n-1), head)
	node := getNode(length-n, head)
	node.Next = nThNode.Next
	return head
}

func read(head *ListNode) []int {
	var res []int
	for {
		if head == nil {
			break
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

func convertToPositive(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func TestRemoveNthFromEnd(t *testing.T) {
	var head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	read(head)
	head = removeNthFromEnd(head, 4)
	read(head)
	t.Log(head)
}

// URL: https://leetcode.com/problems/merge-two-sorted-lists/description/?envType=problem-list-v2&envId=linked-list
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	// merge them
	len1 := lenList(list1)
	last1 := getNode(len1, list1)
	last1.Next = list2

	// sort
	l := lenList(list1)
	arr := read(list1)
	for i := 0; i < l-1; i++ {
		minIndex := i

		for j := i + 1; j < l; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}

	return arrToList(arr)
}

func arrToList(arr []int) *ListNode {
	var head = &ListNode{
		Val: arr[len(arr)-1],
	}
	for i := len(arr) - 2; i >= 0; i-- {
		node := &ListNode{
			Val:  arr[i],
			Next: head,
		}
		head = node
	}

	return head
}

func TestMergeTwoLists(t *testing.T) {
	var l1 = &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
	}
	var l2 = &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 0,
				},
			},
		},
	}

	res := mergeTwoLists(l1, l2)
	t.Log(read(res))
}
