package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	return canJumpto(nums[1:], nums[0])
}

func canJumpto(nums []int, step int) bool {
	if step >= len(nums) {
		return true
	}
	for i := 0; i < step; i++ {
		if canJumpto(nums[i+1:], nums[i]) {
			return true
		}
	}
	return false
}

func canJump2(nums []int) bool {
	max := 0
	for i := 0; i <= max; i++ {
		if i+nums[i] > max {
			max = i + nums[i]
		}
		if max >= len(nums)-1 {
			return true
		}
	}
	return false
}
