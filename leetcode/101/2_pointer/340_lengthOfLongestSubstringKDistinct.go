package __pointer

import "math"

/**
 *  @ClassName:340_lengthOfLongestSubstringKDistinct
 *  @Description:lengthOfLongestSubstringKDistinct
 *  @Author:jackey
 *  @Create:2021/7/22 下午8:28
 */

// 滑动窗口+哈希表 只能是双指针
// 哈希表用来标记字母出现的最后一个位置

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	sLen := len(s)
	if sLen == 0 || k == 0 {
		return 0
	}

	left, right := 0, 0
	table := make(map[byte]int)
	res := math.MinInt64
	for right < sLen {
		table[s[right]] = right
		right++
		if len(table) == k + 1 {
			idx := getMapMinValue340(table)
			delete(table, s[idx])
			left = idx + 1
		}
		res = max340(res, right - left)
	}
	return res
}

func max340(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMapMinValue340(table map[byte]int) int {
	min := math.MaxInt64
	for _, value := range table {
		if value < min {
			min = value
		}
	}
	return min
}