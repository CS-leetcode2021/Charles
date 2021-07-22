package array

/**
 *  @ClassName:spiralOrder
 *  @Description:leetcode 54 螺旋矩阵
 *  @Author:jackey
 *  @Create:2021/5/25 下午3:09
 */

/*
 *  @Description:   思路
 *  @Param:         需要一个用来表示访问的数组，和一个用来标记步数的变量来判断是都是结束，需要四个方位值来判定是否需要转弯
 *  @Return:
 * 	时间复杂度O（mn），空间复杂度O（mn），有辅助矩阵
 */

func spiralOrder(matrix [][]int) []int {
	// 判定无效的数据
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	// 获取长度
	rows, cols := len(matrix), len(matrix[0])

	// 创建辅助数组，用于标识是否被访问
	visited := make([][]int, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]int, cols)
	}

	// path用来记录输出的总个数
	path := rows * cols

	// 创建辅助数组，用来判断方位
	directions := [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
	directionIndex := 0

	// 遍历数据
	r, c := 0, 0
	res := make([]int, path)
	for i := 0; i < path; i++ {
		// 写入结果集
		res[i] = matrix[r][c]
		// 标识当前的数据是否访问过
		visited[r][c] = 1

		// 判断下一步是否应该转向
		nextRow, nextCol := r+directions[directionIndex][0], c+directions[directionIndex][1]
		if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols || visited[nextRow][nextCol] == 1 {
			directionIndex = (directionIndex + 1) % 4
		}
		r += directions[directionIndex][0]
		c += directions[directionIndex][1]

	}
	return res
}

/*
 *  @Description:   一层一层的遍历
 *  @Param:
 *  @Return:
 * 	时间复杂度O（mn），空间复杂度O（1）
 */

func spiralOrder2(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix) == 0 {
		return nil
	}
	rows, cols := len(matrix), len(matrix[0])
	res := []int{}
	left, right := 0, cols-1
	top, bottom := 0, rows-1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		for i := top + 1; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		if left < right && top < bottom {
			for i := right-1; i >left ; i-- {
				res = append(res,matrix[bottom][i])
			}
			for i := bottom; i >top ; i-- {
				res = append(res,matrix[i][left])
			}
		}

		left++
		right--
		top++
		bottom--
	}
	return res
}
