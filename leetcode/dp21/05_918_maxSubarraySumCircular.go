package dp21

/**
 *  @ClassName:05_918_maxSubarraySumCircular
 *  @Description:环形子数组的最大和
 *  @Author:jackey
 *  @Create:2021/8/10 下午3:51
 */

// 最大数组要不就在数组里面，要不就包含头尾
// 前者直接退化为数组的最大子序列和
// 反之，求出中间数据的子数组最小和，然后整个数组的总和减去即可

// TODO...
func maxSubarraySumCircular(nums []int) int {
	m := len(nums)

	if m == 1 {
		return nums[0]
	}

	sum, all := nums[0], nums[0]

	max, min := nums[0], nums[1]

	for i := 1; i < m; i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		max = max918(max, sum)
		all += nums[i]
	}

	sum = nums[1]
	for i := 2; i < m-1; i++ {
		if sum > 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		min = min918(min, sum)
	}

	return max918(max, all-min)
}

func max918(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min918(i, j int) int {
	if i < j {
		return i
	}
	return j
}
