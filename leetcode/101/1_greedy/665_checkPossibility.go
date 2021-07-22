package __greedy

/**
 *  @ClassName:665_checkPossibility
 *  @Description:checkPossibility 非递减序列 不简单
 *  @Author:jackey
 *  @Create:2021/7/22 下午6:15
 */

// 贪心 95/90
func checkPossibility(nums []int) bool {
	count := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if i == 1 || nums[i] >= nums[i-2] {
				nums[i-1] = nums[i]
			}else {
				nums[i] =nums[i-1]
			}
			count++
		}
	}

	return count<= 1
}

