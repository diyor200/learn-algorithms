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
	for i := 0; i < index; i++ {
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
	nThNode := getNode(n, head-1-n)
}
