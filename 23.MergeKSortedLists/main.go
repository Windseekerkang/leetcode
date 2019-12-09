package main

import "fmt"

func main() {
	node := mergeKLists([]*ListNode{
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  6,
				Next: nil,
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

func mergeKLists1(lists []*ListNode) *ListNode {
	headNode := &ListNode{
		Val:  0,
		Next: nil,
	}
	var lastNode = headNode
	for {
		nextNode := getSmallestNodeHead(lists)
		if nextNode == nil {
			break
		}
		lastNode.Next = nextNode
		lastNode = nextNode
	}
	return headNode.Next
}

func getSmallestNodeHead(lists []*ListNode) *ListNode {
	resultindex := -1
	for key, v := range lists {
		if v != nil {
			resultindex = key
			break
		}
	}
	if resultindex == -1 {
		return nil
	}
	result := &ListNode{}
	for index, listHead := range lists {
		if listHead == nil {
			continue
		}
		if listHead.Val < lists[resultindex].Val {
			resultindex = index
		}
	}
	if lists[resultindex] == nil {
		return nil
	} else {
		result = lists[resultindex]
		lists[resultindex] = lists[resultindex].Next
	}
	return result
}

func mergeKLists(lists []*ListNode) *ListNode {
	size := len(lists)
	if size == 0 {
		return nil
	}
	if size == 1 {
		return lists[0]
	}
	list1 := mergeKLists(lists[0 : size/2])
	list2 := mergeKLists(lists[size/2:])
	return mergeTwo(list1, list2)
}

func mergeTwo(list1, list2 *ListNode) *ListNode {
	var head *ListNode
	cur := &head
	for list1 != nil || list2 != nil {
		if list1 == nil {
			*cur = list2
			list2 = list2.Next
			cur = &((*cur).Next)
		} else if list2 == nil {
			*cur = list1
			list1 = list1.Next
			cur = &((*cur).Next)
		} else {
			if list1.Val <= list2.Val {
				*cur = list1
				list1 = list1.Next
				cur = &((*cur).Next)
			} else {
				*cur = list2
				list2 = list2.Next
				cur = &((*cur).Next)
			}
		}
	}
	return head
}
