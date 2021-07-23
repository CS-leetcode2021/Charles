package __sort

/**
 *  @ClassName:5_selectionSort
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/7/23 下午4:29
 */

// 每次选择最大的放在后面
func selectSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		maxIndex := 0
		//寻找最大的一个数，保存索引值
		for j := 1; j < length-i; j++ {
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
		}
		nums[length-i-1], nums[maxIndex] = nums[maxIndex], nums[length-i-1]
	}
}
