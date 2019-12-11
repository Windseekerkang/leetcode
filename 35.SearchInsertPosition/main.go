package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1}, 1))
}

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[l] == target {
			return l
		}
		if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[l] >= target {
		return l
	} else {
		return l + 1
	}
}
