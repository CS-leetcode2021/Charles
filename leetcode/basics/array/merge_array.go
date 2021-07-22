package array

import (
	"sort"
)

/**
 *  @ClassName:merge_array
 *  @Description:leetcode 56 合并区间 无序的集合
 *  @Author:jackey
 *  @Create:2021/5/25 下午6:28
 */

/*
 *  @Description: 合并之后的新区间可能和后面的区间重合，所以需要一次次遍历，一遍得出一个结果
 *  @Param:		 通过sort.Slice先进性排序，然后再逐一比较
 *  @Return:
 */

func mergeIntervals(intervals [][]int) [][]int {
	// Slice sorts the slice x given the provided less function.
	// It panics if x is not a slice.
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	prev := intervals[0]

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]

		if prev[1] < cur[0] {
			res = append(res, prev)
			prev = cur
		}else {
			prev[1] = MaxMerge(prev[1],cur[1])
		}
	}
	res = append(res,prev)
	return res

}

func MaxMerge(i, j int) int {
	if i>j {
		return i
	}
	return j
}