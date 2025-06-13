package queue

import (
	"fmt"
	"testing"
)

// url: https://leetcode.com/problems/implement-stack-using-queues/description/?envType=problem-list-v2&envId=queue

// using 2 queues
//type MyStack struct {
//	q1 []int
//	q2 []int
//}
//
//func Constructor() MyStack {
//	return MyStack{
//		q1: []int{},
//		q2: []int{},
//	}
//}
//
//func (ms *MyStack) Push(x int) {
//	ms.q2 = append(ms.q2, x)
//	for len(ms.q1) > 0 {
//		ms.q2 = append(ms.q2, ms.q1[0])
//		ms.q1 = ms.q1[1:]
//	}
//
//	ms.q1, ms.q2 = ms.q2, []int{}
//}
//
//func (ms *MyStack) Pop() int {
//	res := ms.q1[0]
//	ms.q1 = ms.q1[1:]
//
//	return res
//}
//
//func (ms *MyStack) Top() int {
//	return ms.q1[0]
//}
//
//func (ms *MyStack) Empty() bool {
//	return len(ms.q1) == 0
//}

//type MyStack struct {
//	q []int
//}
//
//func Constructor() MyStack {
//	return MyStack{
//		q: []int{},
//	}
//}
//
//func (ms *MyStack) Push(x int) {
//	ms.q = append(ms.q, x)
//	for i := 0; i < len(ms.q)-1; i++ {
//		ms.q = append(ms.q, ms.q[0])
//		ms.q = ms.q[1:]
//	}
//}
//
//func (ms *MyStack) Pop() int {
//	top := ms.q[0]
//	ms.q = ms.q[1:]
//	return top
//}
//
//func (ms *MyStack) Top() int {
//	return ms.q[0]
//}
//
//func (ms *MyStack) Empty() bool {
//	return len(ms.q) == 0
//}

func TestMyStack(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	//fmt.Println(stack.Top()) // 2
	fmt.Println(stack.Pop()) // 2
	fmt.Println(stack.Empty())
}

type MyQueue struct {
	inStack  []int // used for push
	outStack []int // used for pop/peek
}

func Constructor() MyQueue {
	return MyQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

func (mq *MyQueue) Push(x int) {
	mq.inStack = append(mq.inStack, x)
}

func (mq *MyQueue) Pop() int {
	if len(mq.outStack) == 0 {
		mq.transfer()
	}
	val := mq.outStack[len(mq.outStack)-1]
	mq.outStack = mq.outStack[:len(mq.outStack)-1]
	return val
}

func (mq *MyQueue) Peek() int {
	if len(mq.outStack) == 0 {
		mq.transfer()
	}
	return mq.outStack[len(mq.outStack)-1]
}

func (mq *MyQueue) Empty() bool {
	return len(mq.inStack) == 0 && len(mq.outStack) == 0
}

func (mq *MyQueue) transfer() {
	for len(mq.inStack) > 0 {
		n := len(mq.inStack)
		mq.outStack = append(mq.outStack, mq.inStack[n-1])
		mq.inStack = mq.inStack[:n-1]
	}
}

func TestMyQueue(t *testing.T) {
	//stack := Constructor()
	//stack.Push(1)
	//stack.Push(2)
	//stack.Push(3)
	//stack.Push(4)
	//fmt.Println(stack.Pop()) // 1
	//stack.Push(5)
	//fmt.Println(stack.Pop()) // 1
	//fmt.Println(stack.Pop()) // 1
	//fmt.Println(stack.Pop()) // 1
	//fmt.Println(stack.Pop()) // 1
	////fmt.Println(stack.Peek()) // 1
	////fmt.Println(stack.Pop())  // 1
	//fmt.Println(stack.Empty())
	s := "leetcode"
	fmt.Println(s[2])
}

// url: https://leetcode.com/problems/first-unique-character-in-a-string/?envType=problem-list-v2&envId=queue

func firstUniqChar(s string) int {
	type index struct {
		count int
		index int
	}

	m := make(map[byte]index, len(s))

	for i := range s {
		if val, ok := m[s[i]]; ok {
			val.count++
			m[s[i]] = val
		} else {
			m[s[i]] = index{count: 1, index: i}
		}
	}

	first := len(s)
	for k, _ := range m {
		if m[k].count == 1 && m[k].index < first {
			first = m[k].index
		}
	}

	if first == len(s) {
		return -1
	}
	return first
}

func veryFast(s string) int {
	var letters = [26]int{}

	for i := 0; i < len(s); i++ {
		letters[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if letters[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

func TestFirstUniqChar(t *testing.T) {
	s := "leetcode"
	//s := "aabbb"
	fmt.Println(firstUniqChar(s))
	fmt.Println(veryFast(s))
	fmt.Println('l')
	fmt.Println('a')
}
