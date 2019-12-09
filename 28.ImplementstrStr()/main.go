package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(strStrKMP("a", "a"))
}

func strStrKMP(haystack string, needle string) int {
	fs := *(*[]byte)(unsafe.Pointer(&haystack))
	ns := *(*[]byte)(unsafe.Pointer(&needle))
	i := 0
	j := 0
	next := getNext(ns)
	for i < len(fs) && j < len(ns) {
		if j == -1 || fs[i] == ns[j] {
			i++
			j++
		} else {
			//i = i - j + 1
			//j = 0
			j = next[j]
		}
	}
	if j == len(ns) {
		return i - j
	}
	return -1
}

func getNext(bs []byte) []int {
	next := make([]int, len(bs))
	if len(bs) == 0 {
		return next
	}
	next[0] = -1
	i := 0
	j := -1
	for i < len(bs)-1 {
		if j == -1 || bs[i] == bs[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

func strStrBF(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	size := len(needle)
	for i := 0; i <= len(haystack)-size; i++ {
		if haystack[i:i+size] == needle {
			return i
		}
	}
	return -1
}
