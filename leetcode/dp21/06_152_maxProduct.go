package dp21

/**
 *  @ClassName:06_152_maxProduct
 *  @Description:乘积最大子数组
 *  @Author:jackey
 *  @Create:2021/8/10 下午6:26
 */

// 同53 如果使用前一段的最优子结构转化而来会导致出错
// eg:{5,6,-3,4,-3}
// 按照53题的解法，会得到最大值30
// 正确的值应该是全部的乘积

// 需要维护两个dp数组
// 如果当前是正值，寻找前一个位置的正的最大值最优解
// 如果当前是负值，寻找前一个位置的负的最小值最优解

func maxProduct(nums []int) int {

	m := len(nums)
	if m == 1 {
		return nums[0]
	}
	dpMax := make([]int, m)
	dpMin := make([]int, m)

	dpMax[0], dpMin[0] = nums[0], nums[0]

	for i := 1; i < m; i++ {
		// 当前可能是正的也是负值
		dpMax[i] = Max152(dpMax[i-1]*nums[i], Max152(nums[i], dpMin[i-1]*nums[i]))
		dpMin[i] = Min152(dpMin[i-1]*nums[i], Min152(nums[i], dpMax[i-1]*nums[i]))
	}
	res := dpMax[0]

	for i := 0; i < len(dpMax); i++ {
		res = Max152(res, dpMax[i])
	}

	return res
}

func Max152(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func Min152(i, j int) int {
	if i > j {
		return j
	}
	return i
}

// 优化
// 100/74
func maxProduct2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	maxV, minV := nums[0], nums[0]
	resMax := nums[0]
	for i := 1; i < n; i++ {
		maxV, minV = Max152(maxV*nums[i], Max152(nums[i], nums[i]*minV)), Min152(minV*nums[i], Min152(nums[i], nums[i]*maxV))
		resMax = Max152(resMax, maxV)
	}
	return resMax
}
