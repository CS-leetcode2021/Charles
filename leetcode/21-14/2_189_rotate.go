package _1_14

/**
 *  @ClassName:2_189_rotate
 *  @Description:旋转数组
 *  @Author:jackey
 *  @Create:2021/7/26 上午10:53
 */
// 使用辅助数组
//
func rotate(nums []int, k int) {
	// 创建一个新的数组
	newNums := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		newNums[(i+k)%len(nums)] = nums[i]
	}

	copy(nums, newNums)

}

// 使用原地置换
// 90/77
func rotate2(nums []int, k int) {
	// 先将整个数组翻转一遍
	reverse189(nums)

	// 将前k个数字翻转一遍，后面也要翻转
	reverse189(nums[:k%len(nums)])
	reverse189(nums[k%len(nums):])
}

func reverse189(nums []int) {
	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
