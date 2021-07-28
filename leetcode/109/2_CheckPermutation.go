package _09

import (
	"sort"
	"strings"
)

/**
 *  @ClassName:2_CheckPermutation
 *  @Description:判断字符串是否是重排
 *  @Author:jackey
 *  @Create:2021/7/28 下午2:58
 */

// 100/100 或者是使用map
func CheckPermutation(s1 string, s2 string) bool {
	m,n := len(s1),len(s2)
	if m != n{
		return false
	}

	var ch1 , ch2 [26]int
	for i := 0 ; i < m ; i++{
		ch1[s1[i]-'a']++
		ch2[s2[i]-'a']++
	}

	if ch1 == ch2 {
		return true
	}
	return false
}

// 使用内置函数
// 切割字符串然后排序后在拼接
// 判断是否相同即可
// 100/16
func CheckPermutation2(s1 string, s2 string) bool {
	tmp1 := strings.Split(s1, "")
	tmp2 := strings.Split(s2, "")
	sort.Strings(tmp1)
	res1 := strings.Join(tmp1, "")
	sort.Strings(tmp2)
	res2 := strings.Join(tmp2, "")
	return res1 == res2
}