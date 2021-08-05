package _1_14

/**
 *  @ClassName:11_784_letterCasePermutation
 *  @Description:字母大小写全排列
 *  @Author:jackey
 *  @Create:2021/8/3 下午4:50
 */

// 按照回溯思想写
// 大写变小写、小写变大写：字符 ^= 32 （大写 ^= 32 相当于 +32，小写 ^= 32 相当于 -32）
// 大写变小写、小写变小写：字符 |= 32 （大写 |= 32 就相当于+32，小写 |= 32 不变）
// 大写变大写、小写变大写：字符 &= -33 （大写 ^= -33 不变，小写 ^= -33 相当于 -32）

func letterCasePermutation(S string) []string {
	var (
		ans    []string
		dfs    func(start int, path []byte)
		length = len(S)
		str    = []byte(S)
	)

	inArea := func(b byte) bool {
		return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
	}

	dfs = func(start int, sub []byte) {
		if start == length {
			ans = append(ans, string(sub))
			return
		}
		// 未修改当前字符(字母或者数字)的一条分支
		dfs(start+1, str)
		// 修改当前字母的的另一条分支
		if inArea(str[start]) {
			// 大小写转换
			str[start] ^= 32
			dfs(start+1, str)
		}
	}
	dfs(0, str)
	return ans
}
