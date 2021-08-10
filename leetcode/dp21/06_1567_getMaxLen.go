package dp21

/**
 *  @ClassName:06_1567_getMaxLen
 *  @Description:乘积为正数的最长子数组长度
 *  @Author:jackey
 *  @Create:2021/8/10 下午7:57
 */

// 两个数组positive和negative，其中的元素表示的含义是以下标i为结尾时，当前子序列中存在的连续正数数组/负数数组的长度。

// 100/84
func getMaxLen(nums []int) int {
	m := len(nums)

	if m == 1 {
		if nums[0] > 0 {
			return 1
		} else {
			return 0
		}
	}

	positive := make([]int, m)
	negative := make([]int, m)
	res := 0
	if nums[0] > 0 {
		positive[0], negative[0] = 1, 0
	} else if nums[0] < 0 {
		positive[0], negative[0] = 0, 1
	}// else nums[0] == 0 两个数组都是0

	for i := 1; i < m; i++ {
		if nums[i] > 0 {
			//由于正数×正数=正数，因此仅需在先前乘积为正数的最长子序列长度上加一即可
			positive[i] = positive[i-1] + 1
			if negative[i-1] == 0 {
				negative[i] = 0
			} else {
				negative[i] = negative[i-1] + 1
			}
		}
		if nums[i] == 0 {
			positive[i], negative[i] = 0, 0
		}

		if nums[i] < 0 {
			//由于负数×负数=正数，因此判断是否存在以下标i-1为结尾的乘积为负数的最长子序列
			if negative[i-1] != 0 { // 存在
				positive[i] = negative[i-1] + 1
			} else { // 不存在
				positive[i] = 0
			}
			//由于正数×负数=负数，因此仅需在先前乘积为正数的最长子序列长度上加一即可
			negative[i] = positive[i-1] + 1

		}
		res = Max1567(res, positive[i])

	}
	return res
}

func Max1567(i, j int) int {
	if i > j {
		return i
	}
	return j
}
