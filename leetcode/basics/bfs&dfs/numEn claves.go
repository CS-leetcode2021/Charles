package binary_tree

/**
 *  @ClassName:numEn claves
 *  @Description:leetcode-1020 飞地的数量
 *  @Author:jackey
 *  @Create:2021/7/9 下午7:37
 */
// DFS
// 1、如果触碰到边界则说明是可以离开的
// 2、从四个边界开始逆向DFS，将其修改为2
// 3、统计二维数组是1的个数

var EncX = [4]int{1, -1, 0, 0}
var EncY = [4]int{0, 0, 1, -1}

func numEnclaves(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		dfsClaves(grid,i,0)
		dfsClaves(grid,i,n-1)
	}
	for i := 0; i < n; i++ {
		dfsClaves(grid,0,i)
		dfsClaves(grid,m-1,i)
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
	}
	return count


}

func dfsClaves(grid [][]int, x, y int) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] != 1 {
		return
	}
	grid[x][y] = 2
	for i := 0; i < 4; i++ {
		tmp_x := x + EncX[i]
		tmp_y := y + EncY[i]
		if isInClaves(grid,tmp_x,tmp_y) && grid[tmp_x][tmp_y] == 1 {
			dfsClaves(grid,tmp_x,tmp_y)
		}
	}
}

func isInClaves(grid [][]int, x, y int) bool {
	if  x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return false
	}
	return true
}