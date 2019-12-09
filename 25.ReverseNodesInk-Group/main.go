package main

import "fmt"

func main() {
	node := reverseKGroup(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}, 4)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k <= 1 {
		return head
	}
	realHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	index := 1
	var vhead = realHead
	var vlast = head
	var p = head
	var useP = head
	for p != nil {
		useP = p
		p = p.Next
		useP.Next = vhead.Next
		vhead.Next = useP
		if index%k == 0 {
			vhead = vlast
			vlast = p
		}
		index++
	}
	index -= 1
	if index%k > 0 {
		newHead := &ListNode{
			Val:  0,
			Next: nil,
		}
		p = vhead.Next
		for p != nil {
			useP = p
			p = p.Next
			useP.Next = newHead.Next
			newHead.Next = useP
		}
		vhead.Next = newHead.Next
	}
	return realHead.Next
}
