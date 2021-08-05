package _1_14

/**
 *  @ClassName:6_567_checkInclusion
 *  @Description:字符串的排列,字符串s1的排列之一是s2的子串就可以
 *  @Author:jackey
 *  @Create:2021/7/28 下午1:22
 */

// 创建一个map统计s1中字母出现的个数
// 双指针：统计左右指针之间出现的字母个数是否等于s1的个数，如果相等便返回true
// 100/100
func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}

	var ch1, ch2 [26]int

	for i, ch := range s1 {	// 滑动窗口
		ch1[ch-'a']++
		ch2[s2[i]-'a']++
	}

	if ch1 == ch2 {
		return true
	}

	for i := m; i < n; i++ {
		ch2[s2[i]-'a']++
		ch2[s2[i-m]-'a']--
		if ch1 == ch2 {
			return true
		}
	}
	return false
}

// 官方2.使用一个tag表示s1\s2之间不同值的个数

func checkInclusion2(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for i, ch := range s1 {
		cnt[ch-'a']--
		cnt[s2[i]-'a']++
	}
	diff := 0
	for _, c := range cnt[:] {
		if c != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for i := n; i < m; i++ {
		x, y := s2[i]-'a', s2[i-n]-'a'
		if x == y {
			continue
		}
		if cnt[x] == 0 {
			diff++
		}
		cnt[x]++
		if cnt[x] == 0 {
			diff--
		}
		if cnt[y] == 0 {
			diff++
		}
		cnt[y]--
		if cnt[y] == 0 {
			diff--
		}
		if diff == 0 {
			return true
		}
	}
	return false
}


// 双指针：没太明白

func checkInclusion3(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for _, ch := range s1 {
		cnt[ch-'a']--
	}
	left := 0
	for right, ch := range s2 {
		x := ch - 'a'
		cnt[x]++
		for cnt[x] > 0 {
			cnt[s2[left]-'a']--
			left++
		}
		if right-left+1 == n {
			return true
		}
	}
	return false
}
