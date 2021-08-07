package dp21

/**
 *  @ClassName:3_213_robII
 *  @Description:打家劫舍2
 *  @Author:jackey
 *  @Create:2021/8/7 下午5:43
 */

// 区别 最后一个和第一个相连
// 分成两个子数组进行dp求解，取最大
func rob2(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return Max213(nums[0], nums[1])
	}

	res := Max213(rob(nums[:len(nums)-1]), rob(nums[1:]))
	return res

}

func subRob(nums []int) int {
	tmp1, tmp2 := 0,0

	for _, v := range nums {
		tmp1, tmp2 = tmp2, Max213(tmp2, tmp1+v)
	}
	return tmp2

}

func Max213(i, j int) int {
	if i > j {
		return i
	}
	return j
}
