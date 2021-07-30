package _1

/**
 *  @ClassName:9_994_orangesRotting
 *  @Description:腐烂的橘子
 *  @Author:jackey
 *  @Create:2021/7/30 下午8:48
 */

// 68/41
func orangesRotting(grid [][]int) int {
	// 将坏的橘子放进队列
	m ,n := len(grid),len(grid[0])
	queue := make([][]int, 0)
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				count++
			}
		}
	}

	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	day := 0
	for count > 0 && len(queue) > 0 {
		curLen := len(queue)

		for i := 0; i < curLen; i++ {
			curGrid := queue[0]
			queue = queue[1:]
			for _, v := range direction {
				tmp_x := curGrid[0] + v[0]
				tmp_y := curGrid[1] + v[1]
				if tmp_x >= 0 && tmp_x < m && tmp_y >= 0 && tmp_y < n && grid[tmp_x][tmp_y] == 1 {
					grid[tmp_x][tmp_y] = 2
					count--
					queue = append(queue, []int{tmp_x, tmp_y})
				}
			}
		}
		day++
	}
	if count > 0 {
		return -1
	}
	return day
}

// 100/88
func orangesRotting2(grid [][]int) int {
	res := 0                            // 时间
	queue := []int{}                    // 怀橘子队列
	orange := 0                         // 好橘子个数
	col, row := len(grid[0]), len(grid) // 二维数组的长宽
	// 遍历二维数组,将烂橘子的坐标加入队列
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			// 统计好橘子数量
			if grid[i][j] == 1 {
				orange++
			}
			// 怀橘子入队
			if grid[i][j] == 2 {
				queue = append(queue, i*col+j)	// 这里使用的顺序数列记录
			}
		}
	}
	// 上下左右四个方向
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	// bfs
	for orange > 0 && len(queue) != 0 {
		size := len(queue) // 提前保存队列长度
		for i := 0; i < size; i++ {
			// 出队
			node := queue[0]
			queue = queue[1:]
			c, r := node%col, node/col
			// 怀橘子的上下左右遍历
			for j := 0; j < 4; j++ {
				nr, nc := r+dx[j], c+dy[j]
				if nr >= 0 && nr < row && nc >= 0 && nc < col && grid[nr][nc] == 1 {
					orange--         // 好橘子减一
					grid[nr][nc] = 2 // 好橘子变成坏橘子,然后入队
					queue = append(queue, nr*col+nc)
				}
			}
		}
		// 时间加一
		res++
	}
	if orange != 0 { // 好橘子有剩余
		return -1
	}
	return res
}
