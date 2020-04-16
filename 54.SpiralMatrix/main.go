package main

import "fmt"

func main() {
	var l = [][]int{
		{7},
		{9},
		{6},
	}
	fmt.Println(spiralOrder(l))
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) < 1 {
		return []int{}
	}
	if len(matrix) == 1 {
		return matrix[0]
	}
	top, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	result := make([]int, 0, len(matrix)*len(matrix[0]))
	for top <= down && left <= right {
		result = append(result, matrix[top][left:right+1]...)
		top++
		if top > down {
			break
		}
		for i := top; i <= down; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		for i := right; i >= left; i-- {
			result = append(result, matrix[down][i])
		}
		down--
		for i := down; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
	}
	return result
}
