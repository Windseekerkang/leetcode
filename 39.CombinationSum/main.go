package main

import (
	"fmt"
	"sort"
)

func main() {
	result := combinationSum([]int{2, 3, 5}, 8)
	for _, e := range result {
		for _, value := range e {
			fmt.Print(value)
		}
		fmt.Println()
	}
}

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(candidates)
	combinationSumSingle(&result, candidates, []int{}, 0, target, 0)
	return result
}

func combinationSumSingle(result *[][]int, candidates, path []int, sum, target, index int) {
	if sum == target {
		*result = append(*result, path)
		return
	}
	if sum > target {
		return
	}
	for i := index; i < len(candidates); i++ {
		newCandidates := make([]int, len(path)+1)
		copy(newCandidates[:len(path)], path)
		newCandidates[len(newCandidates)-1] = candidates[i]
		combinationSumSingle(result, candidates, newCandidates, sum+candidates[i], target, i)
	}
}
