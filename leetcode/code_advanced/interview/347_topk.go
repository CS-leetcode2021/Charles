package interview

import "sort"

/**
 *  @ClassName:347_topk
 *  @Description:前k个高频的元素
 *  @Author:jackey
 *  @Create:2021/8/5 下午10:09
 */

func topKFrequent(nums []int, k int) []int {
	if len(nums) < k {
		return nil
	}

	// 创建map
	tMap := make(map[int]int)

	for _, v := range nums {
		tMap[v]++
	}
	tVis := make([][2]int, 0)

	for num, count := range tMap {
		tVis = append(tVis, [2]int{num, count})
	}

	sort.Slice(tVis, func(i, j int) bool {
		return tVis[i][1] > tVis[j][1]
	})

	res := make([]int,0)

	for i := 0; i < k; i++ {
		res = append(res, tVis[i][0])
	}
	return res
}
