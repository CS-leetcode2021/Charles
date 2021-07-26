package __DfsBfs

/**
 *  @ClassName:79_exist
 *  @Description:单次搜索
 *  @Author:jackey
 *  @Create:2021/7/25 下午6:40
 */

// 上次的求解
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs79(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func dfs79(board [][]byte, i int, j int, word string, k int) bool {
	if board[i][j] != word[k] {
		return false
	}

	if k == len(word)-1 {
		return true
	}

	temp := board[i][j]
	board[i][j] = byte(' ')

	if 0 <= i-1 && dfs79(board, i-1, j, word, k+1) {
		return true
	}
	if i+1 < len(board) && dfs79(board, i+1, j, word, k+1) {
		return true
	}
	if 0 <= j-1 && dfs79(board, i, j-1, word, k+1) {
		return true
	}
	if j+1 < len(board[0]) && dfs79(board, i, j+1, word, k+1) {
		return true
	}

	board[i][j] = temp
	return false
}

// dfs递归求解+回溯
// 1，对每一个元素进行遍历
// 2、合适的元素进行递归
// 3、对需要判定的word做出判定，满足返回
// 4、撤出修改标记

var Dx79 = []int{1, -1, 0, 0}
var Dy79 = []int{0, 0, 1, -1}

func exist2(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if backtracking79(board, i, j, 0, word) {
					return true
				}
			}
		}
	}
	return false
}

func backtracking79(board [][]byte, i, j, tag int, word string) bool {
	if !lsInArray79(board, i, j) || board[i][j] != word[tag] {
		return false
	}
	// tmp记录当前变量，为了撤销操作
	if tag == len(word)-1 {
		return true
	}
	tmp := board[i][j]
	board[i][j] = ' '
	for k := 0; k < 4; k++ {
		tmp_x := i + Dx79[k]
		tmp_y := j + Dy79[k]
		if backtracking79(board, tmp_x, tmp_y, tag+1, word){
			return true
		}
	}
	board[i][j] = tmp
	return false
}

func lsInArray79(board [][]byte, i, j int) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}
	return true
}
