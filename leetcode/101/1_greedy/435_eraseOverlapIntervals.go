package __greedy

import "sort"

/**
 *  @ClassName:435_eraseOverlapIntervals
 *  @Description:mid		// 区间不重叠问题
 *  @Author:jackey
 *  @Create:2021/7/22 上午10:18
 */

// 贪心算法：区间不重叠个数
// idea:取件的结尾十分重要，选择的区间结尾越小，则得到的区间个数越多
// 98/100
func eraseOverlapIntervals(intervals [][]int) int {

	// 自定义排序，按照区间的结尾升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			count++
			// 需要移出前一个相交的数据
			intervals[i] = intervals[i-1]
		}
		/*
			或者是保存一个中间的变量
		eg:
			if intervals[i][0] < prev {
				count++
			}else{
				prev = intervals[i][1]
			}

		*/
	}
	return count
}
