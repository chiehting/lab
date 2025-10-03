package main

import (
	"fmt"
	"time"
)

// ListNode is singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// version 2
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	rootlist := &ListNode{}
	newlist := rootlist
	sum, carry := 0, 0

	for {
		if l1 != nil && l2 != nil {
			sum = l1.Val + l2.Val + carry
			l1, l2 = l1.Next, l2.Next
		} else if l1 == nil && l2 != nil {
			sum = l2.Val + carry
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			sum = l1.Val + carry
			l1 = l1.Next
		} else {
			sum = carry
		}

		carry = sum / 10
		newlist.Val = sum % 10

		if carry > 0 || l1 != nil || l2 != nil {
			newlist.Next = &ListNode{}
			newlist = newlist.Next
		} else {
			break
		}
	}

	return rootlist
}

// version 1
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	var newlist []int
// 	var offset int = 0

// 	for {
// 		newlist = append(newlist, (l1.Val+l2.Val+offset)%10)
// 		offset = (l1.Val + l2.Val + offset) / 10
// 		if l1.Next != nil && l2.Next != nil {
// 			l1 = l1.Next
// 			l2 = l2.Next
// 		} else if l1.Next == nil && l2.Next != nil {
// 			l1.Val = 0
// 			l2 = l2.Next
// 		} else if l1.Next != nil && l2.Next == nil {
// 			l1 = l1.Next
// 			l2.Val = 0
// 		} else {
// 			if offset > 0 {
// 				newlist = append(newlist, offset)
// 			}
// 			break
// 		}
// 	}
// 	return newlistlink(newlist)
// }

func newlistlink(vals []int) *ListNode {
	if len(vals) == 1 {
		return &ListNode{Val: vals[0]}
	}
	return &ListNode{Val: vals[0], Next: newlistlink(vals[1:])}
}

func main() {
	// l1 := newlistlink([]int{2, 4, 3})
	// l2 := newlistlink([]int{5, 6, 4})

	// l1 := newlistlink([]int{0})
	// l2 := newlistlink([]int{0})

	l1 := newlistlink([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := newlistlink([]int{9, 9, 9, 9})

	t1 := time.Now()
	ans := addTwoNumbers(l1, l2)
	fmt.Println(t1.Sub(time.Now()))

	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}
