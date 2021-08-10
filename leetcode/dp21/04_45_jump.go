package dp21

/**
 *  @ClassName:04_45_jump
 *  @Description:跳跃游戏2
 *  @Author:jackey
 *  @Create:2021/8/10 下午3:08
 */

// dp走不通啊，
// 状态方程：
func jump(nums []int) int {
	end := 0
	steps := 0
	max_pos := 0
	for i := 0; i < len(nums)-1; i++ {
		max_pos = maxDistance(max_pos,nums[i]+i)
		if i == end {
			end = max_pos
			steps++
		}
	}
	return steps
}


func maxDistance(x,y int) int  {
	if x > y {
		return x
	}

	return y

}