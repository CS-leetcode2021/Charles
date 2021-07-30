package sign

/**
 *  @ClassName:0730_171_titleToNumber
 *  @Description:序列号
 *  @Author:jackey
 *  @Create:2021/7/30 下午6:38
 */

// 100/5
func titleToNumber(columnTitle string) int {

	tmpBit := make([]int, len(columnTitle))

	for i := 0; i < len(columnTitle); i++ {
		tmpBit[i] = int(columnTitle[i]-'A') + 1
	}

	res := 0
	for i := 0; i < len(tmpBit); i++ {
		exp := len(tmpBit) - i - 1
		mult := tmpBit[i]
		tmp := 1
		for exp > 0 {
			tmp *= 26
			exp--
		}
		res += mult * tmp
	}
	return res
}

// 优化
// 100/62
func titleToNumber2(columnTitle string) int {
	res := 0

	for i := 0; i < len(columnTitle); i++ {
		exp := len(columnTitle) - i - 1
		tmp := 1
		for exp > 0 {
			tmp *= 26
			exp--
		}
		res += (int(columnTitle[i]-'A') + 1) * tmp
	}
	return res
}

// 官网
// 100/100
func titleToNumber3(columnTitle string) (number int) {
	for i, multiple := len(columnTitle)-1, 1; i >= 0; i-- {
		k := columnTitle[i] - 'A' + 1
		number += int(k) * multiple
		multiple *= 26
	}
	return
}
