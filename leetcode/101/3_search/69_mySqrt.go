package __search

/**
 *  @ClassName:69_mySqrt
 *  @Description:求算术平方根，
 *  @Author:jackey
 *  @Create:2021/7/23 上午10:23
 */

// 二分法进行求解
// 100/62
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	l, r := 1, x
	var mid, sqrt int

	for l <= r {
		mid = (r-l)/2 + l
		sqrt = x / mid

		if sqrt == mid {
			return sqrt
		} else if sqrt > mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}

// 牛顿迭代法
// 100/100
func mySqrt2(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r)/2
	}
	return r
}
