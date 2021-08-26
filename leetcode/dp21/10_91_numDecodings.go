package dp21

/**
 *  @ClassName:10_91_numDecodings
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/26 下午1:35
 */

// 注意累加
func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1

	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && (s[i-2]-'0')*10+(s[i-1]-'0') <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}
