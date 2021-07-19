package main

import (
	"bytes"
	"sort"
	"strconv"
)

/**
 *  @ClassName:179_largestNumber
 *  @Description:leetcode-179 最大数
 *  @Author:jackey
 *  @Create:2021/7/19 下午7:38
 */

func largestNumber(nums []int) string {
	if nums == nil || len(nums) == 0 {
		return ""
	}

	strs := make([]string, len(nums))

	for i := 0; i < len(nums); i++ {
		strs[i] = strconv.Itoa(nums[i])
	}

	// 自定义对比的规则
	sort.Slice(strs, func(i, j int) bool {
		s12 := strs[i] + strs[j]
		s21 := strs[j] + strs[i]
		// 如果字符串s12大于s21,那么我们希望s12排在s21前
		return s12 > s21
	})	// 已经排序结束了

	if strs[0] == "0" {
		return "0"
	}

	var buffer bytes.Buffer
	for i := range strs {
		buffer.WriteString(strs[i])
	}
	return buffer.String()
}
