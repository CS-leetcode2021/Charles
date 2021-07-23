package __sort

import "sort"

/**
 *  @ClassName:347_topKFrequent
 *  @Description:前 K 个高频元素
 *  @Author:jackey
 *  @Create:2021/7/23 下午7:15
 */

// 没有用桶排序
func topKFrequent(nums []int, k int) []int {

	if len(nums) < k {
		return nil
	}
	tMap := make(map[int]int)

	tVis := make([][2]int,0)
	for i := 0; i < len(nums); i++ {
		if _, ok := tMap[nums[i]]; !ok {
			tMap[nums[i]]=1
			tVis = append(tVis,[2]int{nums[i],0})
		}else {
			tMap[nums[i]]++
		}
	}

	for i := 0; i < len(tVis); i++ {
		tVis[i][1] = tMap[tVis[i][0]]
	}

	sort.Slice(tVis, func(i, j int) bool {
		return tVis[i][1] > tVis[j][1]
	})

	res := []int{}

	for i := 0; i < k; i++ {
		res = append(res,tVis[i][0])
	}

	return res
}