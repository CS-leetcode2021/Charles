package __greedy

import "sort"

/**
 *  @ClassName:452_findMinArrowShots
 *  @Description:Medium 用最少的箭引爆气球	// 区间重叠问题
 *  @Author:jackey
 *  @Create:2021/7/22 下午3:30
 */


func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	count := 1
	tmp := []int{points[0][0],points[0][1]}

	for i := 1; i < len(points); i++ {

		if points[i][0] <= tmp[1]{
			tmp[0] = points[i][0]
			if points[i][1] < points[i-1][1] {
				tmp[1] = points[i][1]
			}
		}else {
			tmp[0] = points[i][0]
			tmp[1] = points[i][1]
			count++
		}
	}
	return count
}


func findMinArrowShots3(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	maxRight := points[0][1]
	ans := 1
	for _, p := range points {
		if p[0] > maxRight {
			maxRight = p[1]
			ans++
		}
	}
	return ans
}