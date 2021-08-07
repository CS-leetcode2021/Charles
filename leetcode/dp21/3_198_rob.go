package dp21

/**
 *  @ClassName:3_198_rob
 *  @Description:打家劫舍
 *  @Author:jackey
 *  @Create:2021/8/7 下午5:42
 */


func Max198(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	tag1, tag2 := nums[0], Max198(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		tag1, tag2 = tag2, Max198(tag2, tag1+nums[i])
	}

	return tag2

}