package binary_tree

/**
 *  @ClassName:closedIsland
 *  @Description:leetcode-1254. 统计封闭岛屿的数目 BFS
 *  @Author:jackey
 *  @Create:2021/7/9 下午8:17
 */
// bfs
// 1、首先是边界上的所有陆地清空，边界上的陆地不可能被包围,同理与边界上的陆地相连的也不可能被包围，使用DFS会简便很多
// 2、遍历，如果是陆地，则进行BFS，每次结束算是一个封闭岛屿

var LX = [4]int{1, -1, 0, 0}
var LY = [4]int{0, 0, 1, -1}

func closedIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			dfsLand(grid,i,0)
		}
		if grid[i][n-1] == 0 {
			dfsLand(grid,i,n-1)
		}
	}

	for i := 0; i < n; i++ {
		if grid[0][i] == 0 {
			dfsLand(grid,0,i)
		}
		if grid[m-1][i] == 0 {
			dfsLand(grid,m-1,i)
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				// 这里使用dfs 实现双百，如果使用bfs会浪费空间内存，毕竟要使用队列等结构
				dfsLand(grid, i, j)
				//bfsLand(grid, i, j)
				count++
			}
		}
	}

	return count
}

func bfsLand(grid [][]int, x, y int) {
	q := new(Queue)
	q.EnQueue(x)
	q.EnQueue(y)
	for q.Size() != 0 {
		cursize := q.Size() / 2
		for i := 0; i < cursize; i++ {
			x = q.Dequeue()
			y = q.Dequeue()
			for k := 0; k < 4; k++ {
				tmp_x := x + LX[k]
				tmp_y := y + LY[k]
				if isInland(grid, tmp_x, tmp_y) && grid[tmp_x][tmp_y] == 0 {
					grid[tmp_x][tmp_y] = 1
					q.EnQueue(tmp_x)
					q.EnQueue(tmp_y)
				}
			}
		}
	}
}

func dfsLand(grid [][]int, x, y int) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] != 0{
		return
	}
	grid[x][y] = 1
	for i := 0; i < 4; i++ {
		tmp_x := x + LX[i]
		tmp_y := y + LY[i]

		if isInland(grid,tmp_x,tmp_y) && grid[tmp_x][tmp_y] == 0 {
			dfsLand(grid,tmp_x,tmp_y)
		}
	}
}
func isInland(grid [][]int, x, y int) bool {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return false
	}
	return true
}
