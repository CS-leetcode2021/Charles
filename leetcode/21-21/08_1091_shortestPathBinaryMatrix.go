package _1_21

/**
 *  @ClassName:08_1091_shortestPathBinaryMatrix
 *  @Description:二进制中的最短路径
 *  @Author:jackey
 *  @Create:2021/8/13 下午8:44
 */

// 如何定义最短路径
func shortestPathBinaryMatrix(grid [][]int) int {
	r := len(grid)
	if grid == nil || r == 0 || grid[0][0] == 1 || grid[r-1][r-1] == 1 {
		return -1
	}

	if len(grid) == 1 && grid[0][0] == 1 {
		return 1
	}
	dir := []int{1, 0, -1}
	grid[0][0] = 1
	queue := make([]int, 0)
	queue = append(queue, 0)

	for len(queue) > 0 {
		x, y := queue[0]/r, queue[0]%r
		queue = queue[1:]
		for _, i := range dir {
			for _, j := range dir {
				if i == j && i == 0 {
					continue
				}

				nx, ny := x+i, y+j

				if nx < 0 || nx >= r || ny < 0 || ny >= r || grid[nx][ny] != 0 {
					continue
				} else {
					queue = append(queue, nx*r+ny)
					grid[nx][ny] = grid[x][y] + 1
					if nx == r-1 && ny == r-1 {
						return grid[nx][ny]
					}
				}
			}
		}
	}
	return -1
}
