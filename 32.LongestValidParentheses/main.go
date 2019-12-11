package main

import "fmt"

//todo DP实现
func main() {
	fmt.Println(longestValidParentheses(")()())()()("))
}

func longestValidParentheses(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if len(s)-i-1 < result {
			return result
		}
		if s[i] == ')' {
			continue
		}
		v := 1
		inresult := 0
		for j := i + 1; j < len(s); j++ {
			if s[j] == ')' {
				v--
			} else {
				v++
			}
			if v < 0 {
				break
			}
			if v == 0 {
				inresult = j - i + 1
			}
		}
		if inresult > result {
			result = inresult
		}
	}
	return result
}
