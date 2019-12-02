package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"abab", "aba", ""}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	defaultSlice := *(*[]byte)(unsafe.Pointer(&strs[0]))
	for _, s := range strs {
		if len(s) == 0 {
			return ""
		}
		if len(defaultSlice) >= len(s) {
			defaultSlice = *(*[]byte)(unsafe.Pointer(&s))
		}
	}
	back := 0
	isOver := false
	for {
		for _, s := range strs {
			bytes := *(*[]byte)(unsafe.Pointer(&s))

			if back >= len(defaultSlice) {
				isOver = true
				break
			}

			if defaultSlice[back] != bytes[back] {
				isOver = true
				break
			}
		}
		if isOver {
			break
		}
		back++
	}
	if back == 0 {
		return ""
	}
	defaultSlice = defaultSlice[0:back]
	return *(*string)(unsafe.Pointer(&(defaultSlice)))
}
