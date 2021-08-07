package _09

/**
 *  @ClassName:4_canPermutePalindrome
 *  @Description:回文排列
 *  @Author:jackey
 *  @Create:2021/8/6 下午4:04
 */

// 使用map
func canPermutePalindrome(s string) bool {
	tmap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		tmap[s[i]]++
	}

	tag := false

	for _, count := range tmap {
		if count&1 == 0 {
			if tag {
				return false
			}
			tag = true
		}
	}
	return true
}
