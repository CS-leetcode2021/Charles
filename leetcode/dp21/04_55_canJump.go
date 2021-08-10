package dp21

/**
 *  @ClassName:04_55_canJump
 *  @Description:跳跃游戏
 *  @Author:jackey
 *  @Create:2021/8/10 下午2:46
 */

// dp[i] 保存当下做能到达的最远位置
// 添加一层判定 line:17
func canJump2(nums []int) bool {

	n := len(nums)
	max := 0
	for i := 0; i < n; i++ {
		if i <= max { // 你不可能跑到比当前位置更远的位置
			max = Max55(max, i+nums[i])
		} // [3,2,1,0,4] 未通过
		if max >= n-1 {
			return true
		}
	}

	return false

}

func Max55(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 倒着来就可以

func canJump(nums []int) bool {
	n := len(nums)
	last := n - 1

	for i := n - 2; i >= 0; i-- {
		if i+nums[i] >= last {
			last = i
		}
	}

	return last == 0
}
