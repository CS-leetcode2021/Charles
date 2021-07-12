package main

/**
 *  @ClassName:16_myPow
 *  @Description:剑指offer-16 数值的整数次方 同leetcode-50
 *  @Author:jackey
 *  @Create:2021/7/12 下午5:56
 */

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	if n < 0 {
		n = -n
		x = 1 / x
	}

	temp := myPow(x, n/2)

	if n%2 == 0 {
		return temp * temp
	}
	return x * temp * temp
}

func myPow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	res := 1.0

	for n > 1 {
		if n&1 == 1 {
			res *= x
			n--
		} else {
			x *= x			// x的平方
			n = n >> 1
		}
	}
	return res*x
}
