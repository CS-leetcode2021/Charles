package binary_tree

// leetcode-733	双百
// 类似于1034边框着色
// 1、只要对应坐标的连通着色问题
// 2、一次DFS找到所有的连通量，改变其颜色为新的颜色
// 3、为了避免重复访问，需要设置tag来标记
// 4、缺点是空间复杂度过高（解决：如果新的颜色等于旧的颜色，直接返回即可）

var Fx = [4]int{1, -1, 0, 0}
var Fy = [4]int{0, 0, 1, -1}

func floodFill(grid [][]int, r0, c0 int, color int) [][]int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return nil
	}
	val := grid[r0][c0]
	//m, n := len(grid), len(grid[0])
	//tag := make([][]int, m)
	//for i := 0; i < m; i++ {
	//	tag[i] = make([]int, n)
	//}
	//dfsflood(grid, tag, r0, c0, val, color)
	if color == val {	// 优化：如果新颜色等于原来的颜色，直接返回即可
		return grid
	}
	dfsflood(grid, r0, c0, val, color)
	return grid
}

func dfsflood(grid [][]int, x, y, val, color int) {
	if !isInflood(grid, x, y) || grid[x][y] != val {
		return
	}
	grid[x][y] = color
	// tag[x][y] = 1
	for i := 0; i < 4; i++ {
		tmp_x := x + Fx[i]
		tmp_y := y + Fy[i]
		if isInflood(grid, tmp_x, tmp_y) && grid[tmp_x][tmp_y] == val{
			dfsflood(grid , tmp_x, tmp_y, val, color)
		}
	}

}

func isInflood(grid [][]int, x, y int) bool {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return false
	}
	return true
}
