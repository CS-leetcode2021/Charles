package __DfsBfs

import "strings"

/**
 *  @ClassName:51_N-queens
 *  @Description:N-皇后问题
 *  @Author:jackey
 *  @Create:2021/7/26 下午5:16
 */

// 每行没列只有一个皇后
// 相邻和斜着也不能有
// dfs+回溯
// 1、创建一个标记数组
// 2、创建当前放置的皇后数量，为n时添加进结果集
// 3、对每一个边界进行判定，这里是上下左右共计8个方向
// 4、添加回撤操作
// 5、只需要对每一行遍历就可以

func solveNQueens(n int) [][]string {
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j, _ := range board {
			board[i][j] = "."
		}
	}

	res := [][]string{}

	return res
}

func backtracking51(board [][]string, r, n int, res *[][]string) {
	if r == n { // 终止条件
		// 复合条件的结果，添加进结果集
		tmp := make([]string, 0)
		for i := 0; i < len(board); i++ {
			s := strings.Join(board[i], "")
			tmp = append(tmp, s)
		}
		*res = append(*res, tmp)
		return
	}

	for i := r; i < n; i++ {
		if isValid(board, r, i, n) { // 选取复合的结果

		}
	}
	// 未完待续......
}

func isValid(board [][]string, r, c int, n int) bool {
	// 可以设置辅助数组记录出现过的皇后，列、对角线，反对角线
	return true
}
