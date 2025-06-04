package leetcode

import (
	"fmt"
	"testing"
)

//func isPowerOfTwo(n int) bool {
//	var res = 1
//	for {
//		if res == n {
//			return true
//		} else if res > n {
//			return false
//		}
//
//		res *= 2
//	}
//}

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func Test_IsPowerOfTwo(t *testing.T) {
	fmt.Println(8 & (8 - 1))
	n := 20
	if isPowerOfTwo(n) == true {
		t.Log("ok")
	} else {
		t.Fatal("not ok")
	}
}

// url: https://leetcode.com/problems/power-of-three/submissions/1642116340/?envType=problem-list-v2&envId=recursion
func isPowerOfThree(n int) bool {
	// 1 = 0001
	// 3 = 0011
	// 9 = 1001
	// 27 =
	for n%3 == 0 {
		n /= 3
	}

	return n == 1
}

func Test_IsPowerOfThree(t *testing.T) {
	n := 28
	if isPowerOfThree(n) {
		t.Log("ok")
	} else {
		t.Fatal("not ok")
	}
}

// url: https://leetcode.com/problems/power-of-four/description/?envType=problem-list-v2&envId=recursion
func isPowerOfFour(n int) bool {
	if n < 1 {
		return false
	}
	for n%4 == 0 {
		n /= 4
	}

	return n == 1
}

func Test_IsPowerOfFour(t *testing.T) {
	n := 257
	if isPowerOfFour(n) {
		t.Log("ok")
	} else {
		t.Fatal("not ok")
	}
}

// url: https://leetcode.com/problems/fibonacci-number/?envType=problem-list-v2&envId=recursion
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var a, b = 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

func Test_Fib(t *testing.T) {
	n := 3
	fmt.Println(fib(n))
}
