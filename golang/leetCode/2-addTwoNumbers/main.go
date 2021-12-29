package main

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var newlist []int
	var offset int = 0

	for {
		newlist = append(newlist, (l1.Val+l2.Val+offset)%10)
		offset = (l1.Val + l2.Val + offset) / 10
		if l1.Next != nil && l2.Next != nil {
			l1 = l1.Next
			l2 = l2.Next
		} else if l1.Next == nil && l2.Next != nil {
			l1.Val = 0
			l2 = l2.Next
		} else if l1.Next != nil && l2.Next == nil {
			l1 = l1.Next
			l2.Val = 0
		} else {
			if offset > 0 {
				newlist = append(newlist, offset)
			}
			break
		}
	}
	return newlistlink(newlist)
}

func newlistlink(vals []int) *ListNode {
	if len(vals) == 1 {
		return &ListNode{Val: vals[0]}
	}
	return &ListNode{Val: vals[0], Next: newlistlink(vals[1:])}
}

func main() {
	l1 := newlistlink([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := newlistlink([]int{9, 9, 9, 9})
	t1 := time.Now()
	fmt.Println(addTwoNumbers(l1, l2))
	fmt.Println(t1.Sub(time.Now()))
}

