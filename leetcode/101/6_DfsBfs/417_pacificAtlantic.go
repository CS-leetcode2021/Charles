package __DfsBfs

/**
 *  @ClassName:417_pacificAtlantic
 *  @Description:太平洋大西洋水流问题
 *  @Author:jackey
 *  @Create:2021/7/24 下午9:33
 */

var Dx417 = []int{1, -1, 0, 0}
var Dy417 = []int{0, 0, 1, -1}

func pacificAtlantic(heights [][]int) [][]int {
	// 获取长宽
	m, n := len(heights), len(heights[0])

	// 创建太平洋,大西洋的标记数组
	visTai := make([][]int, m)
	visDa := make([][]int, m)
	for i := 0; i < m; i++ {
		visTai[i] = make([]int, n)
		visDa[i] = make([]int, n)
	}
	// 访问四个边
	for i := 0; i < m; i++ {
		dfs(heights, i, 0, visTai)
	}
	for i := 0; i < n; i++ {
		dfs(heights, 0, i, visTai)
	}

	for i := 0; i < m; i++ {
		dfs(heights, i, n-1, visDa)
	}
	for i := 0; i < n; i++ {
		dfs(heights, m-1, i, visDa)
	}
	res := make([][]int,0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visTai[i][j] == 1 && visDa[i][j] == 1 {
				res = append(res,[]int{i,j})
			}
		}
	}

	return res
}

// 创建dfs访问
func dfs(heights [][]int, i, j int, vis [][]int) {


	vis[i][j] = 1

	for k := 0; k < 4; k++ {
		tmp_x := i + Dx417[k]
		tmp_y := j + Dy417[k]
		// 注意判定相关的条件。坐标在函数中，且没有被访问过，且值大于当下
		if IsIn417(heights, tmp_x, tmp_y) && vis[tmp_x][tmp_y] != 1 && heights[tmp_x][tmp_y] >= heights[i][j]{
			dfs(heights, tmp_x, tmp_y, vis)
		}
	}
	return
}

// 创建判定边界函数
func IsIn417(nums [][]int, i, j int) bool {
	if i < 0 || j < 0 || i >= len(nums) || j >= len(nums[0]) {
		return false
	}
	return true
}
