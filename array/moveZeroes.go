package array

/**
 *  @ClassName:moveZeroes
 *  @Description:leetcode 283 移动零
 *  @Author:jackey
 *  @Create:2021/5/21 上午11:50
 */
/*
 *  @Description:   利用双指针，实现还是优点繁琐
 *  @Param:
 *  @Return:
 */
func moveZeroes(nums []int) {
	n := len(nums)

	s := 0
	t := 0
	for s < n {
		// 寻找第一个不是0的位置
		if nums[s] != 0 {
			t = s
			break
		}
		s++
	}
	// 当前边界指针归零
	s = 0

	for t < n {
		// 不是0的元素向前替换
		if nums[t] != 0 {
			tmp := nums[s]
			nums[s] = nums[t]
			nums[t] = tmp
			s++
		}
		t++
	}
}

/*
 *  @Description:   还是双指针实现，重构代码
 *  @Param:
 *  @Return:
 */

func moveZeroes2(nums []int)  {

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