package __pointer

/**
 *  @ClassName:680_validPalindrome
 *  @Description:validPalindrome 验证回文字符串 同 167、633 使用双指针
 *  @Author:jackey
 *  @Create:2021/7/22 下午7:20
 */

func validPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return helper(s,i+1,j) || helper(s,i,j-1)
		}
		i++
		j--
	}
	return true
}

func helper(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func validPalindrome2(s string) bool {
	if len(s) <= 1 {
		return true
	}
	tag := false
	for i, j := 0, len(s)-1; i < j; {
		if s[i] == s[j] {
			i++
			j--
		}else if s[i+1] == s[j] {	// 这个删除结果应该与下面一个并行，两者取一个正确就行
			i++
			if tag {
				return false
			}
			tag = true
		}else if s[i] == s[j-1] {
			j--
			if tag  {
				return false
			}
			tag = true
		}else {
			return false
		}

	}
	return true
}