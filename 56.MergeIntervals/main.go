package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{
		{1, 3},
		{0, 3},
	}))
}

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] > intervals[j][0] {
			return false
		}
		return true
	})

	result := make([][]int, 0, len(intervals))
	minl, minr := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		l, r := intervals[i][0], intervals[i][1]

		if l > minr {
			result = append(result, []int{minl, minr})
			minr = r
			minl = l
		}
		if r > minr {
			minr = r
			continue
		}
	}
	result = append(result, []int{minl, minr})
	return result
}
