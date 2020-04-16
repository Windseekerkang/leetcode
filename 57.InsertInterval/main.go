package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(insert([][]int{
		{1, 5},
	}, []int{2, 3}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	ret := make([][]int, 0, len(intervals)+1)
	l, r := newInterval[0], newInterval[1]
	for i := 0; i < len(intervals); i++ {
		if in(intervals[i][0], intervals[i][1], newInterval[0]) && in(intervals[i][0], intervals[i][1], newInterval[1]) {
			l = intervals[i][0]
			r = intervals[i][1]
			continue
		}

		if in(intervals[i][0], intervals[i][1], newInterval[0]) {
			l = intervals[i][0]
			continue
		}

		if in(intervals[i][0], intervals[i][1], newInterval[1]) {
			r = intervals[i][1]
			continue
		}

		if in(newInterval[0], newInterval[1], intervals[i][0]) && in(newInterval[0], newInterval[1], intervals[i][1]) {
			continue
		}
		ret = append(ret, intervals[i])
	}
	ret = append(ret, []int{l, r})

	sort.Slice(ret, func(i, j int) bool {
		if ret[i][0] > ret[j][0] {
			return false
		}
		return true
	})

	return ret
}

func in(l, r, in int) bool {
	return in >= l && in <= r
}
