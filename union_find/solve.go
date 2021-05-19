package union_find

/**
 *  @ClassName:solve
 *  @Description:130 被围绕的区域
 *  @Author:jackey
 *  @Create:2021/4/1 下午3:31
 */

var m, n int

func solve01(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m,n=len(board),len(board[0])
	for i := 0; i < m; i++ {
		dfs(board,i,0)
		dfs(board,i,n-1)
	}
	for i := 0; i < n; i++ {
		dfs(board,0,i)
		dfs(board,m-1,i)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '#' {
				board[i][j]='O'
			}else if board[i][j]=='O'{
				board[i][j]='X'
			}
		}
	}


}

func dfs(board [][]byte, x, y int) {
	if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != 'O' {
		return
	}
	board[x][y]='#'
	dfs(board,x+1,y)
	dfs(board,x-1,y)
	dfs(board,x,y+1)
	dfs(board,x,y-1)
}

func solve(board [][]byte)  {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m,n=len(board),len(board[0])

	count,parent,size := uf(m*n+1)
	dummy := m*n
	for i := 0; i < m; i++ {
		if board[i][0]=='O' {
			union(i*n,dummy,count,parent,size)
		}
		if board[0][n-1]=='O' {
			union(i*n+n-1,dummy,count,parent,size)
		}
	}

	for i := 0; i < n; i++ {
		if board[0][i]=='O' {
			union(i,dummy,count,parent,size)
		}
		if board[m-1][i]=='O' {
			union(n*(m-1)+i,dummy,count,parent,size)
		}
	}

	d := [][]int{{1,0},{0,1},{0,-1},{-1,0}}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' {
				for k := 0; k < 4; k++ {
					x := i+d[k][0]
					y := j+d[k][1]
					if board[x][y] == 'O' {
						union(x*n+y,i*n+j,count,parent,size)
					}
				}
			}
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if connected(dummy, i*n+j, parent) {
				board[i][j]='X'
			}
		}
	}
}





