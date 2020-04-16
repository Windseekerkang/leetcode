package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

//func trap(height []int) int {
//	var water  = 0
//	var xyM  = make(map[int][]int)
//	for x, y := range height {
//		for i := 1; i<=y; i++ {
//			xyM[i] = append(xyM[i],x)
//		}
//	}
//	for _, value := range xyM {
//		if len(value) < 2{
//			continue
//		}
//		water = water + value[len(value)-1]-value[0]-len(value)+1
//	}
//	return water
//}

func trap(height []int) int {
	size := len(height)
	if size < 3 {
		return 0
	}
	hl, hr := make([]int, size), make([]int, size)
	hl[0], hr[size-1] = 0, 0
	hl[1], hr[size-2] = height[0], height[size-1]
	maxl, maxr := height[0], height[size-1]

	for l := 2; l < size; l++ {
		r := size - l - 1
		maxl = Max(height[l-1], maxl)
		maxr = Max(height[r+1], maxr)
		hl[l] = maxl
		hr[r] = maxr
	}

	var water int = 0

	for key, v := range height {
		water += Max(Min(hl[key], hr[key])-v, 0)
	}
	return water
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
