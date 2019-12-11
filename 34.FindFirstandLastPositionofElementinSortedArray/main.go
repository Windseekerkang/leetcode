package main

import "fmt"

func main() {
	fmt.Println(searchRange2([]int{1, 2, 3}, 2))
}

func searchRange(nums []int, target int) []int {
	left, right := 0, 0
	for index, num := range nums {
		if num == target {
			left = index
			if index == len(nums)-1 {
				return []int{index, index}
			} else {
				for index, n := range nums[left+1:] {
					if n != target {
						right = left + index
						return []int{left, right}
					}
				}
				return []int{left, len(nums) - 1}
			}
		}
	}
	return []int{-1, -1}
}

func searchRange2(nums []int, target int) []int {
	left, right := -1, -1
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	left = l
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	} else {
		if left == len(nums)-1 {
			return []int{left, left}
		} else {
			for index, n := range nums[left+1:] {
				if n != target {
					right = left + index
					return []int{left, right}
				}
			}
			return []int{left, len(nums) - 1}
		}
	}
}
