package _1_14

/**
 *  @ClassName:2_283_moveZeroes
 *  @Description:移动零元素
 *  @Author:jackey
 *  @Create:2021/7/26 上午11:30
 */

// 双指针，每次交换
// 也可以倒着来
func moveZeroes(nums []int)  {

	left,right,n := 0,0,len(nums)

	for right < n {
		if nums[right] != 0 {
			tmp := nums[right]
			nums[right] = nums[left]
			nums[left] = tmp
			left++
		}
		right++
	}

}