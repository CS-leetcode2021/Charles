package union_find

import (
	"fmt"
	"go_leetcode/union_find"
)

/**
 *  @ClassName:solve_test
 *  @Description: 130 solve_test.go
 *  @Author:jackey
 *  @Create:2021/4/1 下午9:29
 */
func mainTest() {
	/*leetcode 130 被围绕的区域 测试通过*/
	board := [][]byte{
		{'O', 'O', 'O', 'O', 'X', 'X'}, {'O', 'O', 'O', 'O', 'O', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'O'}, {'O', 'X', 'O', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'O'}, {'O', 'X', 'O', 'O', 'O', 'O'},
	}
	m, n := len(board), len(board[0])
	count, parent, size := union_find.Uf(m*n + 1)

	fmt.Println()
	for i := 0; i < len(parent); i++ {
		fmt.Print(parent[i])
		fmt.Print("\t")
		if (i+1)%6 == 0 {
			fmt.Println()
		}
	}
	fmt.Println("------------")

	dummy := m * n
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			union_find.Union(i*n, dummy, count, parent, size)
		}
		if board[i][n-1] == 'O' {
			union_find.Union(i*n+n-1, dummy, count, parent, size)
		}
	}

	for i := 0; i < n-1; i++ {
		if board[0][i] == 'O' {
			union_find.Union(i, dummy, count, parent, size)
		}
		if board[m-1][i] == 'O' {
			union_find.Union(n*(m-1)+i, dummy, count, parent, size)
		}
	}

	fmt.Println()
	for i := 0; i < len(parent); i++ {
		fmt.Print(parent[i])
		fmt.Print("\t")
		if (i+1)%6 == 0 {
			fmt.Println()
		}
	}
	fmt.Println("------------")

	d := [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' {
				for k := 0; k < 4; k++ {
					x := i + d[k][0]
					y := j + d[k][1]
					if board[x][y] == 'O' {
						union_find.Union(x*n+y, i*n+j, count, parent, size)
					}
				}
			}
		}
	}

	fmt.Println()
	for i := 0; i < len(parent); i++ {
		fmt.Print(parent[i])
		fmt.Print("\t")
		if (i+1)%6 == 0 {
			fmt.Println()
		}
	}
	fmt.Println("------------")

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if !union_find.Connected(dummy, i*n+j, parent) {
				board[i][j] = 'X'
			}
		}
	}

	fmt.Println()
	for i := 0; i < len(parent); i++ {
		fmt.Print(parent[i])
		fmt.Print("\t")
		if (i+1)%6 == 0 {
			fmt.Println()
		}
	}
	fmt.Println("------------")

}

