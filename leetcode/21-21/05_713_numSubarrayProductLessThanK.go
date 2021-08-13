package _1_21

/**
 *  @ClassName:05_713_numSubarrayProductLessThanK
 *  @Description:numSubarrayProductLessThanK
 *  @Author:jackey
 *  @Create:2021/8/13 下午7:06
 */

// 乘积小于k的子数组
// 要求必须是连续的子数组
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	l := 0
	prod := 1
	res := 0
	for r := 0; r < len(nums); r++ {
		prod *= nums[r]

		for prod >= k {
			prod /= nums[l]
			l++

		}
		res += r - l + 1 // 计算新增的数量
	}
	return res

}
