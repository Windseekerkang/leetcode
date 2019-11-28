package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(romanToInt1("MCMXCIV"))
}

func romanToInt(s string) int {
	back := 0
	bytes := *(*[]byte)(unsafe.Pointer(&s))
	for i := 0; i < len(bytes); i++ {
		switch bytes[i] {
		case 'M':
			back += 1000
		case 'D':
			back += 500
		case 'C':
			if i != len(bytes)-1 {
				if bytes[i+1] == 'D' {
					back += 400
					i++
					continue
				}
				if bytes[i+1] == 'M' {
					back += 900
					i++
					continue
				}
			}
			back += 100
		case 'L':
			back += 50
		case 'X':
			if i != len(bytes)-1 {
				if bytes[i+1] == 'L' {
					back += 40
					i++
					continue
				}
				if bytes[i+1] == 'C' {
					back += 90
					i++
					continue
				}
			}
			back += 10
		case 'V':
			back += 5
		case 'I':
			if i != len(bytes)-1 {
				if bytes[i+1] == 'V' {
					back += 4
					i++
					continue
				}
				if bytes[i+1] == 'X' {
					back += 9
					i++
					continue
				}
			}
			back += 1
		}

	}
	return back
}

func romanToInt1(s string) int {
	back := 0
	pre := 0
	bytes := *(*[]byte)(unsafe.Pointer(&s))
	for _, b := range bytes {
		cur := romanToNumberMap[b]
		if pre < cur {
			back -= pre * 2
		}
		back += cur
		pre = cur
	}
	return back
}

var romanToNumberMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}
