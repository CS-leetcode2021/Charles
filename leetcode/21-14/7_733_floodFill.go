package _1_14

/**
 *  @ClassName:7_733_floodFill
 *  @Description:图像渲染
 *  @Author:jackey
 *  @Create:2021/7/28 下午2:16
 */

// leetcode-733
// 类似于1034边框着色
// 1、只要对应坐标的连通着色问题
// 2、一次DFS找到所有的连通量，改变其颜色为新的颜色
// 3、为了避免重复访问，需要设置tag来标记 // 这个可以省略
// 4、缺点是空间复杂度过高

// 96/100
var Fx = [4]int{1, -1, 0, 0}
var Fy = [4]int{0, 0, 1, -1}

func floodFill(grid [][]int, r0, c0 int, color int) [][]int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return nil
	}
	val := grid[r0][c0]

	if color == val {
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
	for i := 0; i < 4; i++ {
		tmp_x := x + Fx[i]
		tmp_y := y + Fy[i]
		dfsflood(grid , tmp_x, tmp_y, val, color)
	}

}

func isInflood(grid [][]int, x, y int) bool {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return false
	}
	return true
}
