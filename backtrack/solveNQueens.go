package backtrack

import "strings"

/**
 *  @ClassName:solveNQueens
 *  @Description:51 leetcode NQueues
 *  @Author:jackey
 *  @Create:2021/4/8 下午1:47
 */

func solveNQueens(n int) [][]string {
	bd := make([][]string, n)

	/*构建二维数组*/
	for i := 0; i < n; i++ {
		bd[i] = make([]string, n)
	}

	/*初始化二维数组*/
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			bd[i][j] = "."
		}
	}

	res := [][]string{}
	helper(0, bd, &res, n)

	return res
}

func helper(row int, bd [][]string, res *[][]string, n int) {
	/*已经到达指定的行，返回结果就好*/
	if row == n {
		temp := make([]string, len(bd))

		for i := 0; i < len(bd); i++ {
			temp[i] = strings.Join(bd[i], "")
		}
		/*每一次结束只是一个结果*/
		*res = append(*res, temp)
		return
	}
	for c := 0; c < n; c++ {
		if isVaild(row, c, n, bd) {
			/*选择*/
			bd[row][c] = "Q"
			/*递归选择*/
			helper(row+1, bd, res, n)
			/*这里相当于回撤*/
			bd[row][c] = "."
		}
	}

}

/*判断是否可用*/
func isVaild(r, c int, n int, bd [][]string) bool {
	for i := 0; i < r; i++ {
		for j := 0; j < n; j++ {
			if bd[i][j] == "Q" && (j == c || i+j == r+c || i-j == r-c) {
				return false
			}
		}
	}
	return true
}
