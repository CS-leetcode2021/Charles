package binary_tree

/**
 *  @ClassName:pacificAtlantic
 *  @Description:leetcode-417. 太平洋大西洋水流问题
 *  @Author:jackey
 *  @Create:2021/7/9 下午1:22
 */

// -----------------------
var dirX = [4]int{-1, 1, 0, 0}
var dirY = [4]int{0, 0, -1, 1}

func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 {
		return nil
	}

	res := [][]int{}
	m, n := len(heights), len(heights[0])
	ansda := make([][]int, m)
	anstai := make([][]int, m)
	for i := 0; i < m; i++ {
		ansda[i] = make([]int, n)
		anstai[i] = make([]int, n)
	}
	var dfsPacificAtlantic func(hegihts,and [][]int,x,y int)
	dfsPacificAtlantic = func (heights, ans [][]int, x, y int) {
		ans[x][y] = 1
		for i := 0; i < 4; i++ {
			tmp_x := x + dirX[i]
			tmp_y := y + dirY[i]

			if judgeIn(heights, tmp_x, tmp_y) && heights[tmp_x][tmp_y] >= heights[x][y] && ans[tmp_x][tmp_y] == 0  {
				dfsPacificAtlantic(heights, ans, tmp_x, tmp_y)
			}
		}
	}

	for i := 0; i < m; i++ {
		dfsPacificAtlantic(heights, anstai, i, 0)
	}
	for i := 0; i < n; i++ {
		dfsPacificAtlantic(heights, anstai, 0, i)
	}
	for i := 0; i < m; i++ {
		dfsPacificAtlantic(heights, ansda, i, n-1)
	}
	for i := 0; i < n; i++ {
		dfsPacificAtlantic(heights, ansda, m-1, i)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ansda[i][j] == 1 && anstai[i][j] == 1 {
				res = append(res,[]int{i,j})
			}
		}
	}
	return res
}

func judgeIn(heights [][]int, i, j int) bool {
	if i < 0 || j < 0 || i >= len(heights) || j >= len(heights[0]) {
		return false
	}
	return true
}

