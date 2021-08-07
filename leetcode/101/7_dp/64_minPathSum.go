package __dp

/**
 *  @ClassName:64_minPathSum
 *  @Description:最小路径和
 *  @Author:jackey
 *  @Create:2021/8/7 下午9:10
 */

// dp[i][j] 表示走到grid[i][j]位置是最小的路径和
// dp[i][j] = min{dp[i-1][j],dp[i][j-1]}+grid[i][j]
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	//遍历

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				res[i][j] = grid[0][0]
			} else if i == 0 {
				res[i][j] = res[i][j-1] + grid[i][j]
			} else if j == 0 {
				res[i][j] = res[i-1][j] + grid[i][j]
			} else {
				res[i][j] = min64(res[i-1][j], res[i][j-1]) + grid[i][j]
			}
		}
	}
	return res[m-1][n-1]
}

func min64(i, j int) int {
	if i > j {
		return j
	}
	return i
}

// 优化空间复杂度

func minPathSum2(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	dp := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[j] = grid[i][j]
			} else if i == 0 {
				dp[j] = dp[j-1] + grid[i][j]
			} else if j == 0 {
				dp[j] = dp[j] + grid[i][j]
			} else {
				dp[j] = min64(dp[j], dp[j-1]) + grid[i][j]
			}
		}
	}
	return dp[n-1]
}

// 再次优化
func minPathSum3(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	dp := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j != 0 {
				dp[j] = dp[j-1]
			} else if i != 0 && j != 0 {
				dp[j] = min64(dp[j], dp[j-1])
			}
			dp[j] += grid[i][j]
		}
	}
	return dp[n-1]
}
