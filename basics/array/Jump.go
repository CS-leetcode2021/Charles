package array

/**
 *  @ClassName:Jump
 *  @Description:leetcode  45 跳跃游戏2
 *  @Author:jackey
 *  @Create:2021/5/19 下午7:29
 */
/*
 *  @Description:   跳跃的次数，逆向贪心算法，寻找最后一步跳跃前所在的位置，该位置能通过跳跃到达最后一个位置
 *  @Param:         数组
 *  @Return:        最少的次数
 */

func Jump(nums []int) int {
	n := len(nums)
	pos := n-1
	steps := 0
	for pos > 0  {
		for i := 0; i < pos; i++ {
			if i+nums[i] >= pos {
				pos = i
				steps++
				break
			}
		}
	}
	return steps
}

/*
 *  @Description:   正向的贪心算法：设置当前位置下的最远边界，更新最大到达的边界
 *  @Param:         数组
 *  @Return:        最少的步数
 */

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
