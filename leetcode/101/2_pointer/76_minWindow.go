package __pointer

/**
 *  @ClassName:76_minWindow
 *  @Description:hard 滑动窗口
 *  @Author:jackey
 *  @Create:2021/7/22 上午11:35
 */

// 84/66
func minWindow(s string, t string) string {
	tmap := map[byte]int{}

	for i := 0; i < len(t); i++ {
		if _, ok := tmap[t[i]]; !ok {
			tmap[t[i]] = 1
		} else {
			tmap[t[i]]++
		}
	}

	// 滑动窗口
	count, l, slen := 0, 0, len(s)+1
	min_start := 0
	min_end := len(s)
	for i := 0; i < len(s); i++ {
		if _, ok := tmap[s[i]]; ok {
			tmap[s[i]]--
			if tmap[s[i]] >= 0 {	// 如果是小于0,就说明多了
				count++
			}
		}

		for count == len(t) {
			if i-l+1 < slen {	// 比之前的小才更新结果
				min_start = l
				min_end = i
				slen = i-l+1
			}

			if _, ok := tmap[s[l]]; ok {
				tmap[s[l]]++
				if tmap[s[l]] > 0 {	// 等于0 代表当前字符端依旧有该字符。不缺
					count--
				}
			}
			l++
		}
	}
	if slen > len(s) {
		return ""
	}

	return s[min_start:min_end+1]

}