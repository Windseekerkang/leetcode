package main

import "fmt"

func main() {
	fmt.Println(firstMissingPositive([]int{-1, 4, 2, 1, 9, 10}))
}

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 && nums[i] < len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}

	for key, value := range nums {
		if value != key+1 {
			return key + 1
		}
	}
	return len(nums) + 1
}
