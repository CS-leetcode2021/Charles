package array

/**
 *  @ClassName:removeDuplicates
 *  @Description:leetcode 26 删除有序数组的重复项
 *  @Author:jackey
 *  @Create:2021/5/20 上午8:58
 */

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[slow] = nums[i]
			slow++
		}
	}
	return slow

}

func removeDuplicates1(nums []int) int{
	if len(nums)==0 {
		return 0
	}

	n := len(nums)
	s,t := 1,1

	for t < n {
		if nums[t] != nums[s-1] {
			nums[s]=nums[t]
			s++
		}
		t++
	}
	return s
}

/*
 *  @Description:   leetcode 80 删除有序数组中的重复项II，是每个元素出现的次数，最多出现两次，每次跳跃两次
 *  @Param:
 *  @Return:
 */

func removeDuplicatesII(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	s, f := 2, 2
	for f < n {
		if nums[f] != nums[s - 2] {
			nums[s] = nums[f]
			s++
		}
		f++
	}
	return s
}