package main

import "fmt"

func main() {

	node := swapPairs(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	})
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	realHead := &ListNode{
		Val:  0,
		Next: head,
	}
	var p = realHead
	var pre = realHead
	index := 0
	for p != nil {
		if index%2 == 0 && index != 0 {
			pre.Next.Next = p.Next
			p.Next = pre.Next
			pre.Next = p
			pre = p.Next
			p = p.Next.Next
		} else {
			p = p.Next
		}
		index++
	}
	return realHead.Next
}
