package main

import "sort"

/**
 *  @ClassName:40_getLeastNumbers
 *  @Description:剑指offer-40 最小的第k个数字
 *  @Author:jackey
 *  @Create:2021/7/19 下午3:21
 */

// 基于快排、大根堆、选择排序
func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k <= 0 {
		return nil
	}
	if k >= len(arr) {
		return arr
	}
	sort.Ints(arr)

	res := arr[:k]
	return res
}
