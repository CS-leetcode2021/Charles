package __greedy

import "sort"

/**
 *  @ClassName:455_assign_cookies
 *  @Description:easy
 *  @Author:jackey
 *  @Create:2021/7/22 上午9:29
 */
// 贪心算法
// 每次寻找复合胃口大小的最小值进行选择
// 两个数组进行排序，递增排序

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return i < j
	})
	sort.Ints(s)
	count := 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(s); j++ {
			if s[j] >= g[i] {
				count++
				s[j] = -1
				break
			}
		}
	}

	return count
}

// 优化
// 双for循环可以进行优化
// 94/94 时间复杂度降低了很多
func findContentChildren2(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	child, cookies := 0, 0
	for child < len(g) && cookies < len(s) {
		if s[cookies] >= g[child] {
			child++
		}
		cookies++
	}

	return child
}
