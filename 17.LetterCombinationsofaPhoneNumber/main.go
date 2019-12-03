package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(letterCombinations(""))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	s := *(*[]byte)(unsafe.Pointer(&digits))
	if len(s) == 1 {
		return bM[s[0]]
	} else {
		ls := letterCombinations(digits[0 : len(s)/2])
		rs := letterCombinations(digits[len(s)/2 : len(s)])
		ret := make([]string, 0, len(ls)*len(rs))
		for _, l := range ls {
			for _, r := range rs {
				ret = append(ret, l+r)
			}
		}
		return ret
	}
}

var bM = map[byte][]string{
	'2': []string{
		"a", "b", "c",
	},
	'3': []string{
		"d", "e", "f",
	},
	'4': []string{
		"g", "h", "i",
	},
	'5': []string{
		"j", "k", "l",
	},
	'6': []string{
		"m", "n", "o",
	},
	'7': []string{
		"p", "q", "r", "s",
	},
	'8': []string{
		"t", "u", "v",
	},
	'9': []string{
		"w", "x", "y", "z",
	},
}
