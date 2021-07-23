package __sort

/**
 *  @ClassName:3_insertionSort
 *  @Description:插入排序
 *  @Author:jackey
 *  @Create:2021/7/23 下午4:15
 */

// 插入排序是双循环
// 顺序从序列中取一个数与左侧的元素们做比较，如果左侧的元素比取的数大，就向右移，直到把取的数插入到不小于左侧元素的位置处。类似于扑克牌的点数排序。

func insertSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i ; j >0 && nums[j] < nums[j-1]; j-- {
			swap(nums,j,j-1)
		}
	}
}