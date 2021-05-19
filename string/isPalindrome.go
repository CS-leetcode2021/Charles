package string

import "strings"

/**
 *  @ClassName:isPalindrome
 *  @Description:125 验证回文串
 *  @Author:jackey
 *  @Create:2021/3/17 上午10:29
 */

/*判断只有自字符串的回文串*/
func isPalindromeS(s string) bool {
	i := 0
	j := len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
/*判断有其他字符的回文串*/
func isPalindrome(s string) bool {
	i := 0
	j := len(s)-1
	s = strings.ToLower(s)

	for ; i < j; {
		if !isString(s[i]){
			i++
		}else if !isString(s[j]){
			j--
		} else if s[i]==s[j] {
			i++
			j--
		}else {
			return false
		}
	}
	return true
}
/*判断是都是字符，参数是byte类型*/
func isString(ch byte) bool {
	return (ch>= 'A'&&ch<='Z')||(ch >= 'a'&& ch<= 'z')||(ch>= '0'&& ch<= '9')
}