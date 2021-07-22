package __pointer

/**
 *  @ClassName:524_findLongestWord
 *  @Description:findLongestWord
 *  @Author:jackey
 *  @Create:2021/7/22 下午8:07
 */

// 将字典里面的字符串一个一个扔到s中去判断是不是子串
// 所谓字典序，就按照整数的比较法则进行判断就可以了，无需额外的考虑
func findLongestWord(s string, dictionary []string) string {
	if len(s) == 0 {
		return ""
	}
	r := ""

	for i := 0; i < len(dictionary); i++ {
		if IsSubstring(s, dictionary[i]) {
			if len(dictionary[i]) > len(r) || (len(dictionary[i]) == len(r) && dictionary[i] < r) {
				r = dictionary[i]
			}
		}
	}
	return r
}

func IsSubstring(s, t string) bool {
	i, j := 0, 0
	for i < len(s) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			i++
		}
		if j == len(t) {
			return true
		}
	}
	return false
}
