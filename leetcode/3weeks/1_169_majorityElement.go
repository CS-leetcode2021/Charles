package _weeks

/**
 *  @ClassName:1_169_majorityElement
 *  @Description:多数元素
 *  @Author:jackey
 *  @Create:2021/8/7 下午7:40
 */

func majorityElement(nums []int) int {
	m := len(nums)
	tag := 1
	res := nums[0]

	for i := 1; i < m; i++ {
		if nums[i] == res {
			tag++
		}else if tag == 0{
			res = nums[i]
			tag = 1
		}else {
			tag--
		}
	}

	count := 0

	for i := 0; i < m; i++ {
		if nums[i] == res {
			count++
		}
	}

	if count > m/2 {
		return res
	}
	return -1


}
