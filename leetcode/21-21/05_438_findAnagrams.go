package _1_21

/**
 *  @ClassName:05_438_findAnagrams
 *  @Description:字母异位词
 *  @Author:jackey
 *  @Create:2021/8/13 下午6:30
 */

// 滑动窗口
// 条件: 字符一致，数量一致
//
func findAnagrams(s string, p string) []int {
	sArr := [26]byte{}
	pArr := [26]byte{}

	sLen, pLen := len(s), len(p)

	if sLen < pLen {
		return nil
	}

	for i := 0; i < pLen; i++ {
		sArr[s[i]-'a']++
		pArr[p[i]-'a']++
	}
	res := make([]int, 0)
	// 固定的是可以比较的
	if sArr == pArr {
		res = append(res, 0)
	}

	l, r := 0, pLen
	for r < sLen {
		sArr[s[r]-'a']++

		//if r-l+1 > pLen {
		//	sArr[s[l]-'a']--
		//	l++
		//	if sArr == pArr {
		//		res = append(res, l)
		//	}
		//}
		// 这样也是对的
		sArr[s[l]-'a']--
		l++
		if sArr == pArr {
			res = append(res,l)
		}
		r++
	}

	return res

}
