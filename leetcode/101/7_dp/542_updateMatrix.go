package __dp

/**
 *  @ClassName:542_updateMatrix
 *  @Description:01矩阵
 *  @Author:jackey
 *  @Create:2021/8/7 下午9:35
 */

// 86/26 bfs操作
// 但是对于一个大小 O(mn) 的二维数组，对每个位置进行四向搜索，最坏情况的时间复
// 杂度（即全是 1）会达到恐怖的 O(m 2 n 2 )。
func updateMatrix2(matrix [][]int) [][]int {

	n, m := len(matrix), len(matrix[0])
	queue := make([][]int, 0)
	for i := 0; i < n; i++ { // 把0全部存进队列，后面从队列中取出来，判断每个访问过的节点的上下左右，直到所有的节点都被访问过为止。
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				point := []int{i, j}
				queue = append(queue, point)
			} else {
				matrix[i][j] = -1
			}
		}
	}
	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 { // 这里就是 BFS 模板操作了。
		point := queue[0]
		queue = queue[1:]
		for _, v := range direction {
			x := point[0] + v[0]
			y := point[1] + v[1]
			if x >= 0 && x < n && y >= 0 && y < m && matrix[x][y] == -1 {
				matrix[x][y] = matrix[point[0]][point[1]] + 1
				queue = append(queue, []int{x, y})
			}
		}
	}

	return matrix
}

func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}

	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 100000
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果本身就是零不需要更改
			if matrix[i][j] == 0 {
				dp[i][j] = 0
			} else {
				if j > 0 {
					dp[i][j] = min542(dp[i][j], dp[i][j-1]+1)
				}
				if i > 0 {
					dp[i][j] = min542(dp[i][j], dp[i-1][j]+1)
				}
			}
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// 如果本身就是零不需要更改
			if matrix[i][j] != 0 {
				if j < n-1 {
					dp[i][j] = min542(dp[i][j], dp[i][j+1]+1)
				}
				if i < m-1 {
					dp[i][j] = min542(dp[i][j], dp[i+1][j]+1)
				}
			}
		}
	}
	return dp
}

func min542(i, j int) int {
	if i > j {
		return j
	}
	return i
}
