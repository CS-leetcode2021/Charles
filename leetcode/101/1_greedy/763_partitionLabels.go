package __greedy

/**
 *  @ClassName:763_partitionLabels
 *  @Description:划分字符串问题
 *  @Author:jackey
 *  @Create:2021/7/22 下午3:59
 */

// 先统计一遍所有字母出现的频率

// 暴力求解，复杂度都比较高
func partitionLabels(s string) []int {
	tmap := make(map[byte]int, 0)

	// 统计字母频率
	for i := 0; i < len(s); i++ {
		if _, ok := tmap[s[i]]; !ok {
			tmap[s[i]] = 1
		} else {
			tmap[s[i]]++
		}
	}

	// 当前访问到的字母标识数组
	currVisit := []byte{}
	// 当前访问到的字母统计map
	currCount := make(map[byte]int, 0)
	// 返回的结果数组
	res := []int{}
	tag := 0
	for i := 0; i < len(s); i++ {
		if _, ok := currCount[s[i]]; !ok {
			currCount[s[i]] = 1
			currVisit = append(currVisit, s[i])
		} else {
			currCount[s[i]]++
		}

		if currCount[s[i]] == tmap[s[i]] {
			sum := 0
			for j := tag; j < len(currVisit); j++ {
				if currCount[currVisit[j]] == tmap[currVisit[j]] {
					sum += currCount[currVisit[j]]
					if j == len(currVisit)-1 {
						res = append(res, sum)
						tag = j + 1
					}
					continue
				} else {
					break
				}

			}
		}
	}
	return res
}

// 优化求解：区间划分
// 寻找每个字母第一次出现和最后一次出现的下标
// 100/100
func partitionLabels2(s string) []int {
	// 记录最后一次出现的位置
	lastPos := make([]int, 26)

	for i, str := range s {
		lastPos[str-'a'] = i
	}

	start, end := 0, 0
	res := []int{}
	for i, str := range s {
		if lastPos[str-'a'] > end {
			end = lastPos[str-'a']
		}

		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return res
}
