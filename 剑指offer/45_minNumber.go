package main

import (
	"bytes"
	"sort"
	"strconv"
)

/**
 *  @ClassName:45_minNumber
 *  @Description:剑指offer-45
 *  @Author:jackey
 *  @Create:2021/7/19 下午5:25
 */

func minNumber(nums []int) string {
	if nums == nil || len(nums) == 0 {
		return ""
	}

	strs := make([]string,len(nums))

	for i := 0; i < len(nums); i++ {
		strs[i] = strconv.Itoa(nums[i])
	}

	sort.Slice(strs, func(i, j int) bool {
		s12 := strs[i] + strs[j]
		s21 := strs[j] + strs[i]	// 升序
		return s12 < s21
	})

	var buffer bytes.Buffer
	for _, v := range strs {
		buffer.WriteString(v)
	}

	return buffer.String()
}