package main

import (
	"bytes"
)

var n int

func main() {
	solveNQueens(4)
}

func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	in := make([]int, n)
	queen(in, 0, &result)
	return result
}

func queen(a []int, cur int, result *[][]string) {
	if cur == len(a) {
		s := make([]string, len(a))
		bufer := bytes.NewBuffer([]byte{})
		for key, value := range a {
			bufer.Reset()
			for i := 0; i < len(a); i++ {
				if i == key {
					bufer.WriteByte('Q')
				} else {
					bufer.WriteByte('.')
				}
			}
			s[value] = bufer.String()
		}
		*result = append(*result, s)
		return
	}
	for i := 0; i < len(a); i++ {
		a[cur] = i
		flag := true
		for j := 0; j < cur; j++ {
			ab := i - a[j]
			temp := 0
			if ab > 0 {
				temp = ab
			} else {
				temp = -ab
			}
			if a[j] == i || temp == cur-j {
				flag = false
				break
			}
		}
		if flag {
			queen(a, cur+1, result)
		}
	}
}
