package __pointer

import "math"

/**
 *  @ClassName:633_judgeSquareSum
 *  @Description:judgeSquareSum easy 类似于两数之和 167
 *  @Author:jackey
 *  @Create:2021/7/22 下午7:06
 */

// 双指针，一左一右滑动
// 100/100
func judgeSquareSum(c int) bool {
	if c == 0 {
		return true
	}
	for i, j := 0, int(math.Sqrt(float64(c))); i < j; {
		if i*i+j*j == c {
			return true
		} else if i*i+j*j < c {
			i--
		} else {
			j--
		}
	}
	return false
}
