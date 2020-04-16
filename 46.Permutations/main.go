package main

import "fmt"

func main() {
	fmt.Println(permute([]int{2, 3, 1}))
}

func permute(nums []int) [][]int {
	var result [][]int
	var path []int
	visited := make([]bool, len(nums))
	helper(&result, path, nums, visited)

	return result
}

func helper(result *[][]int, path []int, nums []int, visited []bool) {
	if len(path) == len(nums) {
		copyOfNums := make([]int, len(path))
		copy(copyOfNums, path)
		*result = append(*result, copyOfNums)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !visited[i] {
			path = append(path, nums[i])
			visited[i] = true
			helper(result, path, nums, visited)
			path = append(path[:len(path)-1], path[len(path):]...)
			visited[i] = false
		}
	}

	return
}
