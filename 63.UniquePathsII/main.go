package main

import "fmt"

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m < 1 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n < 1 {
		return 0
	}
	if 1 == obstacleGrid[0][0] {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {

			switch {
			case i == 0 && j == 0:
				dp[i][j] = 1
			case i == 0 && j != 0 && obstacleGrid[i][j] == 0:
				dp[i][j] = dp[i][j-1]
			case i != 0 && j == 0 && obstacleGrid[i][j] == 0:
				dp[i][j] = dp[i-1][j]
			case i != 0 && j != 0 && obstacleGrid[i][j] == 0:
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}
	return dp[m-1][n-1]
}
