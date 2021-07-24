package __DfsBfs

/**
 *  @ClassName:695_maxAreaOfIsland
 *  @Description:最大的岛屿面积
 *  @Author:jackey
 *  @Create:2021/7/24 下午7:54
 */

var Dx295 = [4]int{1, -1, 0, 0}
var Dy295 = [4]int{0, 0, 1, -1}

func maxAreaOfIsland(grid [][]int) int {

	m,n := len(grid),len(grid[0])

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area := dfs695(grid,i,j)
				res = Max695(res,area)
			}
		}
	}
	return res
}

func dfs695(grid [][]int, i, j int) int {
	if !isInLand(grid, i, j) || grid[i][j] == 0 {
		return 0
	}

	area := 1
	grid[i][j] = 0

	for k := 0; k < 4; k++ {
		tmp_x := i + Dx295[k]
		tmp_y := j + Dy295[k]
		area += dfs695(grid, tmp_x, tmp_y)
	}

	return area

}

func isInLand(grid [][]int, i, j int) bool {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return false
	}
	return true
}

func Max695(i, j int) int {
	if i > j {
		return i
	}
	return j
}