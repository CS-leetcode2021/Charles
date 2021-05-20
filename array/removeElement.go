package array

/**
 *  @ClassName:removeElement
 *  @Description:leetcode 27 移出元素
 *  @Author:jackey
 *  @Create:2021/5/20 下午8:37
 */
/*
 *  @Description:   双指针，一个用于定位有效长度，一个用于遍历数据；先寻找第一个不是val的位置，遍历指针等于该位置，有效长度指针归零，进行遍历替换是val的值
 *  @Param:         数组
 *  @Return:        新数组的有效长度
 */
func removeElement(nums []int, val int) int {
	n := len(nums)
	s := 0
	for s < n {
		if nums[s] == val {
			s++
		}else{
			break
		}
	}

	t:= s
	s = 0

	for t < n {
		if nums[t]!=val {
			nums[s] = nums[t]
			s++
		}
		t++
	}


	return s
}