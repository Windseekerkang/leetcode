package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}, 3))
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix) == 1 {
		return isINSlice(matrix[0], target)
	}

	min := len(matrix) / 2

	if matrix[min][0] > target {
		return searchMatrix(matrix[:min], target)
	}
	return searchMatrix(matrix[min:], target)
}

func isINSlice(s []int, target int) bool {
	if len(s) == 0 {
		return false
	}

	if len(s) == 1 {
		return s[0] == target
	}
	min := len(s) / 2

	if s[min] > target {
		return isINSlice(s[:min], target)
	}
	return isINSlice(s[min:], target)
}
