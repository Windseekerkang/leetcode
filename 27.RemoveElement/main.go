package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums, 3))
}

func removeElement(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[i], nums[count] = nums[count], nums[i]
			count++
		}
	}
	fmt.Println(nums[:count])
	return count
}
