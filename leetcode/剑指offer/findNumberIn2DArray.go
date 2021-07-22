package main

import (
	"sort"
)

/**
 *  @ClassName:lengthOfLongstSubstring
 *  @Description:剑指offer 04 二维数组中查找 // leetcode 240
 *  @Author:jackey
 */

// 空间复杂度太高，O(mn)
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return  false
	}
	m := len(matrix)
	n := len(matrix[0])
	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return  false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == target {
				return  true
			}
		}
	}
	return false
}

// 利用内置函数实现，在数组切片中寻找某个特定的值，如果查到则返回插入值，如果查不到则插入，同时返回插入的位置
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	for _, num := range matrix {
		i := sort.SearchInts(num,target)
		if i< len(num) && num[i] == target {
			return  true
		}
	}
	return false
}

// 优化空间，通过利用递增关系来进行快速查找
// 从左下角开始查找，同样的通过其他角落也可以
func findNumberIn2DArray3(matrix [][]int, target int) bool {
	m := len(matrix)-1
	n := 0

	for m > -1 {
		if n < len(matrix[m]) {
			if matrix[m][n] > target {
				m--
			}else if matrix[m][n] < target{
				n++
			}else if target == matrix[m][n] {
				return true
			}
		}else {
			return false
		}
	}
	return false
}















