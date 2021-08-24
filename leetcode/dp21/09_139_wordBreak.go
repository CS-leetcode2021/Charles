package dp21

/**
 *  @ClassName:08_139_wordBreak
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/24 下午1:03
 */

// map记录相关的单词
// 遍历字符串s，依次判断
// error
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	tMap := make(map[string]int)

	for i := 0; i < len(wordDict); i++ {
		tMap[wordDict[i]]++
	}
	res := make([]byte, 0)
	tag := false
	for i := 0; i < n; i++ {
		res = append(res, s[i])
		// aaaa,aaa
		if _, ok := tMap[string(res)]; ok {
			tag = true
			res = res[0:0]
		} else {
			tag = false
		}
	}
	return tag
}

// dp
// dp[i] 表示s的前i位是否可以用words中的单词表示
// dp[0] = true
// 0<= j < i,dp[j] == true && s[j:i] == true ====>dp[i] == true

func wordBreak2(s string, words []string) bool {
	l := len(s)
	tMap := make(map[string]bool)
	for _, v := range words {
		tMap[v] = true
	}
	dp := make([]bool, l+1)
	dp[0] = true

	for i := 1; i <= l; i++ {
		for j := i - 1; j >= 0; j-- {
			suffix := s[j:i]
			if tMap[suffix] && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[l]
}
