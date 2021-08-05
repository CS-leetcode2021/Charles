package _1_14

import "strings"

/**
 *  @ClassName:4_557_reverseWords
 *  @Description:反转字符串-单词
 *  @Author:jackey
 *  @Create:2021/7/26 上午11:37
 */

// 13/5
func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	res := ""
	for i := 0; i < len(strs); i++ {
		res += string(reverse557(strs[i]))
		if i != len(strs)-1 {
			res += " "
		}
	}
	return res
}

func reverse557(s string) []byte{
	res := []byte{}
	for i := len(s)-1; i >= 0 ; i-- {
		res = append(res,s[i])
	}
	return res
}

// 官方
// 93/37
// 双指针，遍历每一个单次，进行翻转
func reverseWords2(s string) string {
	length := len(s)
	ret := []byte{}
	for i := 0; i < length; {
		start := i
		for i < length && s[i] != ' ' {
			i++
		}
		for p := start; p < i; p++ {
			ret = append(ret, s[start + i - 1 - p])
		}
		for i < length && s[i] == ' ' {
			i++
			ret = append(ret, ' ')
		}
	}
	return string(ret)
}
