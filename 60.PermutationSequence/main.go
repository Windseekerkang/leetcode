package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getPermutation(4, 9))
}

func getPermutation(n int, k int) string {
	var fac = []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
	vals := make([]int, 0, n)
	result := make([]byte, 0, n)
	for i := 1; i <= n; i++ {
		vals = append(vals, i)
	}
	k--
	for i := n; i > 0; i-- {
		r := k % fac[i-1]
		t := k / fac[i-1]
		k = r
		sort.Ints(vals)
		result = append(result, byte('0'+vals[t]))
		vals = append(vals[:t], vals[t+1:]...)
	}
	return string(result)
}
