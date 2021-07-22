package binary_tree

import "sort"

/**
 *  @ClassName:maxAreaOfIsland
 *  @Description:leetcode-695 岛屿的最大面积 BFS
 *  @Author:jackey
 *  @Create:2021/7/8 下午7:32
 */
var tag[][]int

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	tag := make([][]int, m)
	for i := 0; i < m; i++ {
		tag[i] = make([]int, n)
	}
	res := []int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				//tag[i][j] = 1
				p := &Pos{i, j, grid[i][j]}
				num := bfs(grid, p, tag, m, n)
				res = append(res, num)
			}
		}
	}
	sort.Ints(res)
	return res[len(res)-1]
}

func bfs(grid [][]int, pos *Pos,tag [][]int, m, n int) int {
	num := 0
	q := new(QueueP)
	q.EnQueue(pos)
	for q.Size() != 0 {
		curSize := q.Size()
		for i := 0; i < curSize; i++ {
			val := q.Dequeue()
			if tag[val.posX][val.posY] == 1 {
				continue
			}
			tag[val.posX][val.posY] = 1
			num++
			if val.posX-1 >= 0 && val.posX+1 <= m {
				if grid[val.posX-1][val.posY] == 1 && tag[val.posX-1][val.posY] == 0 {
					p := &Pos{val.posX - 1, val.posY, grid[val.posX-1][val.posY]}
					q.EnQueue(p)
				}
				if grid[val.posX+1][val.posY] == 1 && tag[val.posX+1][val.posY] == 0 {
					p := &Pos{val.posX + 1, val.posY, grid[val.posX+1][val.posY]}
					q.EnQueue(p)
				}
			}
			if val.posY-1 >= 0 && val.posY+1 <= n {
				if grid[val.posX][val.posY-1] == 1 && tag[val.posX][val.posY-1] == 0 {
					p := &Pos{val.posX, val.posY - 1, grid[val.posX][val.posY-1]}
					q.EnQueue(p)
				}
				if grid[val.posX][val.posY+1] == 1 && tag[val.posX][val.posY+1] == 0 {
					p := &Pos{val.posX, val.posY + 1, grid[val.posX][val.posY+1]}
					q.EnQueue(p)
				}
			}
		}

	}
	return num
}

// -----------------------------------------
// dfs
func maxAreaOfIsland2(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	area := -1

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				tmp := dfs(grid,i,j)
				area = Max(area,tmp)
			}
		}
	}

	return area
}

func dfs(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	area := 1
	grid[i][j] = 0
	area += dfs(grid,i-1,j)
	area += dfs(grid,i+1,j)
	area += dfs(grid,i,j-1)
	area += dfs(grid,i,j+1)

	return area
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// -----------------------------------------
// other bfs

func maxAreaOfIsland3(grid [][]int) (result int) {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 {
		return 0
	}
	var bfs func(i, j int) int
	bfs = func(i, j int) (count int) {
		arr := [][2]int{{i, j}}
		for len(arr) != 0 {
			a, b := arr[0][0], arr[0][1]
			arr = arr[1:]
			if a >= 0 && a < n && b >= 0 && b < m && grid[a][b] == 1 {
				count++
				grid[a][b] = 0
				arr = append(arr, [][2]int{{a - 1, b}, {a, b - 1}, {a + 1, b}, {a, b + 1}}...)
			}
		}
		return
	}
	for i := range grid {
		for j := range grid[0] {
			result = Max2(result, bfs(i, j))
		}
	}
	return
}

//Max return the maximum number
func Max2(a ...int) int {
	max := a[0]
	for _, v := range a[1:] {
		if v > max {
			max = v
		}
	}
	return max
}










































