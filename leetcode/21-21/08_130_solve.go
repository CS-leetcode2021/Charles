package _1_21

/**
 *  @ClassName:08_130_solve
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/13 下午9:02
 */
var Dx130 = []int{1, -1, 0, 0}
var Dy130 = []int{0, 0, 1, -1}

func solve(board [][]byte) {
	m, n := len(board), len(board[0])

	// 边界遍历
	for i := 0; i < m; i++ {
		dfs130(board, i, 0)
		dfs130(board, i, n-1)
	}
	for i := 0; i < n; i++ {
		dfs130(board, 0, i)
		dfs130(board, m-1, i)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '*' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs130(board [][]byte, i, j int) {
	if !lsIn130(board, i, j) || board[i][j] != 'O' {
		return
	}
	board[i][j] = '*'
	for k := 0; k < 4; k++ {
		tmp_x := i + Dx130[k]
		tmp_y := j + Dy130[k]
		dfs130(board, tmp_x, tmp_y)
	}
	return
}

func lsIn130(board [][]byte, i, j int) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}
	return true
}