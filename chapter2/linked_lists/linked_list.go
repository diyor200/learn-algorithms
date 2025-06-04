package main

import "fmt"

func main() {
	var list = &List{
		Data: 10,
		Next: &List{
			Data: 20,
			Next: &List{
				Data: 30,
				Next: &List{
					Data: 40,
				},
			},
		},
	}

	fmt.Println("Elements of Linked list:", Read(list))
}

type List struct {
	Data int
	Next *List
}

func Read(head *List) []int {
	var res []int
	for {
		res = append(res, head.Data)
		if head.Next == nil {
			break
		} else {
			head = head.Next
		}
	}

	return res
}
