package _1

/**
 *  @ClassName:4_344_reverseString
 *  @Description:反转字符串
 *  @Author:jackey
 *  @Create:2021/7/26 上午11:33
 */

func reverseString(s []byte) {
	m, n := 0, len(s)-1

	for m < n {
		s[m], s[n] = s[n], s[m]
		m++
		n--
	}
}
