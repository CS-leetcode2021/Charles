package _1_21

import "sort"

/**
 *  @ClassName:1_74_searchMatrix
 *  @Description:搜索二维矩阵
 *  @Author:jackey
 *  @Create:2021/8/5 下午7:43
 */

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	low, high := 0, m*n-1
	for low <= high {
		mid := (low + high) >> 1
		tmpr, tmpc := mid/n, mid%n

		if matrix[tmpr][tmpc] == target {
			return true
		} else if matrix[tmpr][tmpc] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return false
}

// 先试用二分查找到target所在的行，在使用二分查找找到target所在的列
func searchMatrix2(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)-1
	// 使用二分查找找到target所在的行
	for l <= r {
		mid := (l + r) / 2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	// 避免一种特殊情况
	// 当matrix的长度为1时,r为-1
	// 这种处理方法确实很low,但是我想不到还有什么好的处理方法啊😥
	if r == -1 {
		r = 0
	}
	left, right := 0, len(matrix[r])-1
	for left <= right {
		mid := (left + right) / 2
		if matrix[r][mid] == target {
			return true
		} else if matrix[r][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// 利用内置函数
func searchMatrix3(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix), func(i int) bool { return matrix[i][0] > target }) - 1
	if row < 0 {
		return false
	}
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}
