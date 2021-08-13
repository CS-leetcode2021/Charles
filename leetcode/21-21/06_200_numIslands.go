package _1_21

/**
 *  @ClassName:06_200_numIslands
 *  @Description: 岛屿数量
 *  @Author:jackey
 *  @Create:2021/8/13 下午7:45
 */

// dfs
// 100/54
var Dx200 = [4]int{1, -1, 0, 0}
var Dy200 = [4]int{0, 0, 1, -1}

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfsLands(grid, i, j)
				res++
			}
		}
	}
	return res

}

func dfsLands(grid [][]byte, x, y int) {

	grid[x][y] = '0'
	for i := 0; i < 4; i++ {
		tmp_x := x + Dx200[i]
		tmp_y := y + Dy200[i]

		if lsInLands(grid, tmp_x, tmp_y) && grid[tmp_x][tmp_y] == '1' {
			dfsLands(grid, tmp_x, tmp_y)
		}
	}
	return
}

func lsInLands(grid [][]byte, x, y int) bool {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return false
	}
	return true
}
