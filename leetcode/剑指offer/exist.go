package main

/**
 *  @ClassName:exist
 *  @Description:offer-12 矩阵中的路径 && leetcode-79 : dfs解法
 *  @Author:jackey
 *  @Create:2021/7/6 下午8:46
 */

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i int, j int, word string, k int) bool {
	if board[i][j] != word[k] {
		return false
	}

	if k == len(word)-1 {
		return true
	}

	temp := board[i][j]
	board[i][j] = byte(' ')

	if 0 <= i-1 && dfs(board, i-1, j, word, k+1)  {
		return true
	}
	if i+1 < len(board) && dfs(board, i+1, j, word, k+1) {
		return true
	}
	if 0 <= j-1 && dfs(board, i, j-1, word, k+1) {
		return true
	}
	if j+1 < len(board[0]) && dfs(board, i, j+1, word, k+1) {
		return true
	}

	board[i][j] = temp
	return false
}
