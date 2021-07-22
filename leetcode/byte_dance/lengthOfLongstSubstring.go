package main

import (
	"strings"
)

/**
 *  @ClassName:lengthOfLongstSubstring
 *  @Description:字节 3 无重复字符的最长子串 滑动窗口解决
 *  @Author:jackey
 *  @Create:2021/6/21 下午7:37
 */

func length(s string) int {
	// 记录每个字符是否出现过
	memo := map[byte]int{}
	n := len(s)
	// 右指针初始化为-1的下标索引
	// ans代表结果
	rp, res := 0, 0
	// i就是左指针边界
	for lp := 0; lp < n; lp++ {
		// 左指针移动需要删除之前的第一个字符，因为已经出现相同字符了
		if lp != 0 {
			delete(memo, s[lp-1])
		}
		// 右指针移动，并且map中没有出现，则代表是没有重复字符出现的
		for rp < n && memo[s[rp]] == 0 {
			memo[s[rp]]++
			rp++
		}
		// 保存结果的最大值
		res = maxString(res, rp-lp)
	}
	return res
}

func maxString(x, y int) int {
	if x < y {
		return y
	}

	return x
}

// 双100% 窗口并不减小，只增不减
// 两个指针表示窗口大小，遍历以此字符串，窗口在遍历过程中滑动或增大
// 1、窗口内没有重复字符：此时判断i+1和end的关系，超过表示遍历到窗口之外了，增大窗口的大小
// 2、窗口内出现重复字符：左右两个指针都增大，滑动窗口位置到重复字符的后一位
// 3、遍历结束，返回end-start的大小即可
func lenght2(s string) int {
	n := len(s)
	lp , rp := 0,0
	for i := 0; i < n; i++ {
		index := strings.Index(s[lp:i],string(s[i]))

		if index == -1 {
			if i+1 > rp {
				rp = i+1
			}
		}else {
			lp += index+1
			rp += index+1
		}
	}

	return rp-lp
}
// 如果需要返回具体的字符串，只需要在窗口增大的时候记录start的指针即可