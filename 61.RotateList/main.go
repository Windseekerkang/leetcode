package main

import "fmt"

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	listNode := rotateRight(l, 2)

	for listNode != nil {
		fmt.Print(listNode.Val, "->")
		listNode = listNode.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	listNode := head
	length := 0
	var endNode *ListNode
	for listNode != nil {
		if listNode.Next == nil {
			endNode = listNode
		}
		listNode = listNode.Next
		length++
	}

	use := k % length
	if use == 0 {
		return head
	}

	tu := length - use
	var rehead *ListNode
	listNode = head
	for i := 1; i <= tu; i++ {
		if i == tu {
			rehead = listNode.Next
			listNode.Next = nil
			endNode.Next = head
		}
		listNode = listNode.Next
	}
	return rehead
}
