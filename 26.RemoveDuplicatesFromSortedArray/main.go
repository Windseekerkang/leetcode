package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicates(nums []int) int {
	nums = remove(nums)
	return len(nums)
}

func remove(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return nums[:i+1]
}
