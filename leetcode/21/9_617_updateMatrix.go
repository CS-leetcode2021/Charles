package _1

/**
 *  @ClassName:9_617_updateMatrix
 *  @Description:01矩阵
 *  @Author:jackey
 *  @Create:2021/7/30 下午7:28
 */
// BFS 需要借助队列实现
// 超内存了 第48个案例没通过
// 添加标记数组依旧超时 41 案例没通过，48通过
var Dx617 = []int{1, -1, 0, 0}
var Dy617 = []int{0, 0, 1, -1}

func updateMatrix(mat [][]int) [][]int {
	// 只需要遍历二维数组为0的项
	m, n := len(mat), len(mat[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != 0 {
				bfsUpdateMat617(mat, i, j)
			}
		}
	}

	return mat
}

func bfsUpdateMat617(mat [][]int, x, y int) {
	q := new(QueueCoordinate)
	q.EnQueue(&Coordinate{x, y, mat[x][y]})

	vis := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		vis[i] = make([]int, len(mat[0]))
	}
	s := -1
	for q.Size() > 0 {
		s++
		curLen := q.Size()
		tag := false
		for i := 0; i < curLen; i++ {
			curCoord := q.DeQueue()
			vis[curCoord.X][curCoord.Y] = 1
			if curCoord.Val == 0 {
				tag = true
				break
			} else {
				for k := 0; k < 4; k++ {
					tmp_x := curCoord.X + Dx617[k]
					tmp_y := curCoord.Y + Dy617[k]
					if isInMat617(mat, tmp_x, tmp_y) && vis[tmp_x][tmp_y] != 1 {
						q.EnQueue(&Coordinate{tmp_x, tmp_y, mat[tmp_x][tmp_y]})
					}
				}
			}
		}
		if tag {
			break
		}
	}

	mat[x][y] = s
}

func isInMat617(mat [][]int, i, j int) bool {
	if i < 0 || j < 0 || i >= len(mat) || j >= len(mat[0]) {
		return false
	}
	return true
}

// 86/26
func updateMatrix2(matrix [][]int) [][]int {

	n, m := len(matrix), len(matrix[0])
	queue := make([][]int, 0)
	for i := 0; i < n; i++ {    // 把0全部存进队列，后面从队列中取出来，判断每个访问过的节点的上下左右，直到所有的节点都被访问过为止。
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				point := []int{i, j}
				queue = append(queue, point)
			} else {
				matrix[i][j] = -1
			}
		}
	}
	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {  // 这里就是 BFS 模板操作了。
		point := queue[0]
		queue = queue[1:]
		for _, v := range direction {
			x := point[0] + v[0]
			y := point[1] + v[1]
			if x >= 0 && x < n && y >= 0 && y < m && matrix[x][y] == -1 {
				matrix[x][y] = matrix[point[0]][point[1]] + 1
				queue = append(queue, []int{x, y})
			}
		}
	}

	return matrix
}