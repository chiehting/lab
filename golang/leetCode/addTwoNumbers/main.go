package main

import "fmt"

/**
 *  https://leetcode.com/problems/add-two-numbers/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode strucka
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbersBest(l1 *ListNode, l2 *ListNode) *ListNode {
	head, flag := l1, 0
	for ; ; l1, l2 = l1.Next, l2.Next {
		l1.Val += l2.Val + flag
		flag = l1.Val / 10
		l1.Val %= 10

		if l1.Next == nil {
			if flag > 0 {
				l1.Next = &ListNode{0, nil}
			} else {
				l1.Next = l2.Next
				break
			}
		}

		if l2.Next == nil {
			if flag > 0 {
				l2.Next = &ListNode{0, nil}
			} else {
				break
			}
		}
	}

	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l1Current, l2Current, soultion, tmp *ListNode
	var less, plus int
	l1Current = l1
	l2Current = l2

	// var solution *ListNode
	for {
		less = (l1Current.Val + l2Current.Val + plus) % 10
		if soultion == nil {
			soultion = newlinklist(less)
		} else if soultion.Next == nil {
			tmp = nextlinklist(less, soultion)
		} else {
			tmp = nextlinklist(less, tmp)
		}

		plus = (l1Current.Val + l2Current.Val + plus) / 10
		if l1Current.Next == nil && l2Current.Next == nil {
			if plus > 0 {
				if soultion.Next == nil {
					tmp = nextlinklist(plus, soultion)
				} else {
					tmp = nextlinklist(plus, tmp)
				}
			}
			break
		}
		if l2Current.Next != nil {
			l2Current = l2Current.Next
		} else {
			l2Current.Val = 0
		}

		if l1Current.Next != nil {
			l1Current = l1Current.Next
		} else {
			l1Current.Val = 0
		}
	}
	return soultion
}

func newlinklist(v int) (newLinklist *ListNode) {
	newLinklist = &ListNode{Val: v}
	return
}

func nextlinklist(v int, superior *ListNode) (newLinklist *ListNode) {
	newLinklist = newlinklist(v)
	superior.Next = newLinklist
	return
}

func main() {
	l1Head := newlinklist(5)

	l2Head := newlinklist(5)

	// addTwoNumbers(l1Head, l2Head)
	addTwoNumbersBest(l1Head, l2Head)
	fmt.Println("ok")
}
