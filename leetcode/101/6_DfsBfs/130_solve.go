package __DfsBfs

/**
 *  @ClassName:130_solve
 *  @Description:被围绕的区域
 *  @Author:jackey
 *  @Create:2021/7/26 下午8:46
 */

// 首先遍历四个边界，将与边界相连的o全部更改为‘*’
// 再对全表进行遍历，将‘O’ 改成‘X’，将‘*’改成‘O’
// 91/97
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
