package main

/**
 *  @ClassName:29_spiralOrder
 *  @Description:剑指offer-29 同 leetcode 54
 *  @Author:jackey
 *  @Create:2021/7/14 下午8:17
 */

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	m,n := len(matrix),len(matrix[0])

	res := make([]int,0)

	left,right := 0,n-1
	top,down := 0,m-1

	for left <= right && top <= down {
		for i := left; i < right; i++ {
			res = append(res,matrix[top][i])
		}

		for i := top; i <= down; i++ {
			res = append(res,matrix[i][right])
		}

		if left <  right && top < down {
			for i := right-1; i >= left; i-- {
				res = append(res,matrix[down][i])
			}
			for i := down-1; i > top; i-- {
				res = append(res,matrix[i][left])
			}
		}
		left++
		right--
		top++
		down--
	}
	return res


}