package _1_21

/**
 *  @ClassName:08_797_allPathsSourceTarget
 *  @Description:图中的所有路径
 *  @Author:jackey
 *  @Create:2021/8/13 下午9:04
 */

// 空间复杂度太高
func allPathsSourceTarget(graph [][]int) [][]int {

	// 先转化为矩阵
	n := len(graph)
	if n == 0 {
		return nil
	}

	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for _, v := range graph[i] {
			grid[i][v] = 1
		}
	}

	res := make([][]int, 0)
	path := []int{0}
	backtracking797(grid, path, 0, &res)

	return res

}

// 回溯

func backtracking797(grid [][]int, path []int, n int, res *[][]int) {
	if n == len(grid)-1 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(grid[n]); i++ {
		if grid[n][i] == 1 {
			path = append(path, i)
			backtracking797(grid, path, i, res)
			path = path[:len(path)-1]
		}
	}
	return

}

// 优化
// 98/37
func allPathsSourceTarget2(graph [][]int) [][]int {

	res := make([][]int, 0)
	path := []int{0}
	backtracking797II(graph, path, 0, &res)
	return res

}

// 回溯

func backtracking797II(grid [][]int, path []int, n int, res *[][]int) {
	if n == len(grid)-1 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(grid[n]); i++ {
		node := grid[n][i]
		path = append(path, node)
		backtracking797(grid, path, node, res)
		path = path[:len(path)-1]
	}
	return

}
