package main

import "fmt"

func main() {
	in := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(in)
	fmt.Println(in)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	if n == 1 {
		return
	}
	reverse(matrix)

	for i := 0; i < n; i++ {
		for j := i + 1; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return
}

func reverse(matrix [][]int) {
	l, r := 0, len(matrix)-1
	for l < r {
		matrix[l], matrix[r] = matrix[r], matrix[l]
		l++
		r--
	}
}
