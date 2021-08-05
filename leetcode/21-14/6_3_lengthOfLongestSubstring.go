package _1_14

import "strings"

/**
 *  @ClassName:6_3_lengthOfLongestSubstring
 *  @Description:无重复的最长子串
 *  @Author:jackey
 *  @Create:2021/7/28 下午12:19
 */
// 使用滑动窗口
// 两个指针初始化都是0
// 需要辅助map来记录当前的字段是否存在相同的字符串
// 向后滑动，每次遇到相同的进行处理，左指针移动
// 100/43	(使用了map，增加了空间复杂度)
func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0

	tMap := make(map[byte]int) // 需要记录出现的坐标
	max := 0
	for i := 0; i < len(s); i++ {
		if _, ok := tMap[s[i]]; !ok {
			tMap[s[i]] = i
			r++
		} else {                 // 当前字符已经存在过，需要移动做左指针，该左表加1
			if tMap[s[i]] >= l { // 该判断是避免"abba"这样的情况，因为先更新b的指针，后更新a的指针，结果导致左指针向右移动，范围扩大
				l = tMap[s[i]] + 1
			}
			tMap[s[i]] = i
			r = i + 1
		}
		max = MaxSubString3(r-l, max)

	}
	return max
}

func MaxSubString3(i, j int) int {
	if i < j {
		return j
	}
	return i
}

// 利用strings.Index来判断当前字符是否出现过，返回对应的左表，并移动双指针进行计算
// 90/99
func lengthOfLongestSubstring2(s string) int {
	r, l := 0, 0
	var ret int
	for i := range s {
		index := strings.Index(s[l:i], string(s[i]))	// 每次选择的是[l:i]区间，不会考虑左侧区间
		if index == -1 {
			r++
		} else {
			r = i + 1
			l += index + 1
		}
		ret = MaxSubString3(len(s[l:r]), ret)
	}
	return ret
}
