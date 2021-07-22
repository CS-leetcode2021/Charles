package binary_tree

// leetcode-1034. 边框着色  DFS 733的强化类型题目
// 有偏差，题目理解没到位

var Cx = [4]int{1, -1, 0, 0}
var Cy = [4]int{0, 0, 1, -1}

func colorBorder(grid [][]int, r0, c0 int, color int) [][]int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return nil
	}
	val := grid[r0][c0]

	//for i := 0; i < len(grid); i++ {
	//	for j := 0; j < len(grid[0]); j++ {
	//		if grid[i][j] == val {
	//			dfsColor(grid,i,j,val,color)
	//		}
	//	}
	//}
	dfsColor(grid,r0,c0,val,color)
	return grid
}

func dfsColor(grid [][]int, x, y, val, color int) {
	if !isInColor(grid, x, y) || grid[x][y] != val {
		return
	}
	grid[x][y] = color

	for i := 0; i < 4; i++ {
		tmp_x := x + Cx[i]
		tmp_y := y + Cy[i]
		if isInColor(grid, tmp_x, tmp_y) && grid[tmp_x][tmp_y] == val {
			dfsColor(grid,tmp_x,tmp_y,val,color)
		}
	}

}

func isInColor(grid [][]int, x, y int) bool {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return false
	}
	return true
}

// 正确解答
func colorBorder2(grid [][]int, r0 int, c0 int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	visited := make([]bool, m*n)
	srcColor := grid[r0][c0]
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// 返回值: 0表示这个点不属于这个联通分量
		if i < 0 || i > m-1 || j < 0 || j > n-1 {
			return 0
		}
		// 因为后面修改了点的颜色属性，所以先要进行visited判断，如果遍历过，那么肯定属于这个连通分量的
		if visited[i*n+j] == true {
			return 1
		}
		//没有被遍历过，而且颜色和点击点不同，那么不属于这个连通分量
		if grid[i][j] != srcColor {
			return 0
		}

		visited[i*n+j] = true
		res := dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1)
		if res < 4 {
			grid[i][j] = color
		}
		return 1
	}
	dfs(r0, c0)
	return grid
}
