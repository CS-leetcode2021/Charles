package array

/**
 *  @ClassName:canJump
 *  @Description:leetcode 55 跳跃游戏
 *  @Author:jackey
 *  @Create:2021/5/19 下午1:18
 */

/**
 * 实时维护可以到达的最远位置，当前位置的最远位置是 i+nums[i]
 */
func canJump(nums []int) bool {
	n := len(nums)

	if n == 0 {
		return false
	}

	if nums[0]==0 {
		return false
	}
	if n == 1 {
		return true
	}


	zero := -1

	for i := n - 2; i >= 0; i-- {

		if zero >0 {
			if nums[i] > zero-i {
				zero = -1
			}
			continue
		}

		if nums[i]== 0{
			zero = i
			continue
		}
	}

	if zero < 0 {
		return true
	}

	return false
}

func canJump2(nums []int) bool {
	// 当前的所能到达的最远边界
	end := 0
	// 到达的最远距离的数组索引
	max_Distance := 0

	for i := 0;i< len(nums);i++	{
		max_Distance = maxDistance(max_Distance,i+nums[i])
		if i == end {
			end = max_Distance
		}
	}
	return end>=len(nums)-1

}

func maxDistance(x,y int) int  {
	if x > y {
		return x
	}

	return y

}

func canJump3(nums []int) bool {
	reach := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		if i > reach {
			return false
		}
		reach = maxDistance(reach,i+nums[i])
	}
	return true
}

func canJump4(nums []int) bool {
	n := len(nums)
	last := n-1

	for i := n - 2; i >= 0; i-- {
		if i+nums[i] >= last {
			last = i
		}
	}

	return last==0
}