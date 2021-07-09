package binary_tree

/**
 *  @ClassName:updateBoard
 *  @Description:leetcode-529 扫雷游戏 DFS
 *  @Author:jackey
 *  @Create:2021/7/9 下午7:11
 */

var Bx = [8]int{1, 1, 1, 0, 0, -1, -1, -1}
var By = [8]int{1, 0, -1, 1, -1, 1, 0, -1}

func updateBoard(board [][]byte, click []int) [][]byte {
	x, y := click[0], click[1]
	if board[x][y] == 'M' { // 点中地雷 游戏结束
		board[x][y] = 'X'
	} else { // 返回新的面板信息
		dfsBoard(board,x,y)
	}

	return board
}

func dfsBoard(board [][]byte, x, y int) {
	count := 0
	for i := 0; i < 8; i++ {
		tmp_x := x + Bx[i]
		tmp_y := y + By[i]
		if tmp_x < 0 || tmp_y < 0 || tmp_x >= len(board) || tmp_y >= len(board[0]) {
			continue
		}
		if board[tmp_x][tmp_y] == 'M' {
			count++
		}
	}

	if count > 0 {
		board[x][y] = byte('0' + count)
	} else {
		board[x][y] = 'B'
		for i := 0; i < 8; i++ {
			tmp_x := x + Bx[i]
			tmp_y := y + By[i]
			if tmp_x < 0 || tmp_y < 0 || tmp_x >= len(board) || tmp_y >= len(board[0]) || board[tmp_x][tmp_y] != 'E' {
				continue
			}
			dfsBoard(board,tmp_x,tmp_y)
		}
	}
}
