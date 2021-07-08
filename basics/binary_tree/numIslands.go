package binary_tree

/**
 *  @ClassName:numIslands
 *  @Description:leetcode-200 岛屿数量 DFS
 *  @Author:jackey
 *  @Create:2021/7/8 下午9:17
 */

// ------------------------------------
// dfsLands
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfsLands(grid, i, j)
				count++
			}
		}
	}
	return count
}

func dfsLands(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	dfsLands(grid, i-1, j)
	dfsLands(grid, i+1, j)
	dfsLands(grid, i, j-1)
	dfsLands(grid, i, j+1)
	return
}

// ----------------------------------------
// bfsLands

func numIslands2(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				bfsLands(grid, i, j)
				count++
			}
		}
	}
	return count
}

func bfsLands(grid [][]byte, i, j int) {
	q := new(Queue)
	q.EnQueue(i)
	q.EnQueue(j)
	for q.Size() != 0 {
		cursize := q.Size() / 2
		for k := 0; k < cursize; k++ {
			i, j = q.Dequeue(), q.Dequeue()
			if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
				continue
			}
			grid[i][j] = '0'
			q.EnQueue(i - 1)
			q.EnQueue(j)
			q.EnQueue(i + 1)
			q.EnQueue(j)
			q.EnQueue(i)
			q.EnQueue(j - 1)
			q.EnQueue(i)
			q.EnQueue(j + 1)

		}
	}
}

// 关于方位的写法

var dx = [4]int{-1, 1, 0, 0}
var dy = [4]int{0, 0, -1, 1}

func dir(i, j int) int{
	x := 0
	for m := 0; m < 4; m++ {
		tmp_i := i + dx[m]
		tmp_j := j + dx[m]
		x = tmp_i + tmp_j
	}
	return x
}
