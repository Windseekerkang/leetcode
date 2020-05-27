package main

import "fmt"

func main() {
	fmt.Println(minPathSum([][]int{
		{1, 3, 2},
		{1, 5, 1},
		{4, 1, 1},
	}))
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m < 1 {
		return 0
	}
	n := len(grid[0])
	if n < 1 {
		return 0
	}

	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch {
			case i == 0 && j == 0:
				dp[j] = grid[i][j]
			case i != 0 && j == 0:
				dp[j] = grid[i][j] + dp[j]
			case i == 0 && j != 0:
				dp[j] = grid[i][j] + dp[j-1]
			case i != 0 && j != 0:
				dp[j] = grid[i][j] + min(dp[j-1], dp[j])
			}
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
