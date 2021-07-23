package __sort

/**
 *  @ClassName:4_bubbleSort
 *  @Description:冒泡排序
 *  @Author:jackey
 *  @Create:2021/7/23 下午4:24
 */


// 两两比较然后交换
func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				//交换
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}