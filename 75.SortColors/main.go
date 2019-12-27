package main

import "fmt"

func main() {
	sortColors([]int{1, 0, 2})
}

func sortColors(nums []int) {
	i, l, r := 0, 0, len(nums)-1
	for i < r {
		switch nums[i] {
		case 2:
			nums[i], nums[r] = nums[r], nums[i]
			r--
		case 1:
			i++
		default:
			nums[i], nums[l] = nums[l], nums[i]
			l++
			i++
		}
	}
	fmt.Printf("%+v", nums)
}
