package binary_tree

/**
 *  @ClassName:solve
 *  @Description:leetcode-130 被围绕的区域 DFS
 *  @Author:jackey
 *  @Create:2021/7/9 下午3:16
 */

var Sx = [4]int{1, -1, 0, 0}
var Sy = [4]int{0, 0, 1, -1}

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		dfsSolve(board, i, 0)
		dfsSolve(board, i, n-1)
	}

	for i := 0; i < n; i++ {
		dfsSolve(board, 0, i)
		dfsSolve(board, m-1, i)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfsSolve(board [][]byte, x, y int) {
	if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) || board[x][y] != 'O' {
		return
	}
	board[x][y] = '#'
	for i := 0; i < 4; i++ {
		tmp_x := x + Sx[i]
		tmp_y := y + Sy[i]

		dfsSolve(board, tmp_x, tmp_y)
	}
	return
}
