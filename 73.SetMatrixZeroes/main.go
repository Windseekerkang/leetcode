package main

import "fmt"

func main() {
	m := [][]int{
		[]int{1, 1, 1},
		[]int{0, 1, 2},
	}
	setZeroes(m)
	fmt.Println(m)
}

type Point struct {
	x int
	y int
}

func setZeroes1(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	zeroIndex := make([]*Point, 0, 1)
	for x, ys := range matrix {
		for y, yi := range ys {
			if yi == 0 {
				zeroIndex = append(zeroIndex, &Point{
					x: x,
					y: y,
				})
			}
		}
	}
	for _, p := range zeroIndex {
		setZero(matrix, p)
	}
	return
}

func setZero(matrix [][]int, p *Point) {
	m := len(matrix[0])
	n := len(matrix)

	for i := 0; i < m; i++ {
		matrix[p.x][i] = 0
	}
	for i := 0; i < n; i++ {
		matrix[i][p.y] = 0
	}
}

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	iscol := false

	for x, ys := range matrix {
		if matrix[x][0] == 0 {
			iscol = true
		}

		for y := 1; y < len(ys); y++ {
			if matrix[x][y] == 0 {
				matrix[x][0] = 0
				matrix[0][y] = 0
			}
		}
	}

	for x := 1; x < len(matrix); x++ {
		for y := 1; y < len(matrix[0]); y++ {
			if matrix[x][0] == 0 || matrix[0][y] == 0 {
				matrix[x][y] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}

	if iscol {
		for i, _ := range matrix {
			matrix[i][0] = 0
		}
	}

	return
}
