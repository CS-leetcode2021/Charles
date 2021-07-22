package sort

/**
 *  @ClassName:quickSort
 *  @Description:快排
 *  @Author:jackey
 *  @Create:2021/7/19 下午6:22
 */

// 快排是一种分而治之的思想，主要使用分治法
// 基准
// 左侧寻找比基准大的数字，后侧寻找比基准小的数字，然后交换
// 双向指针接触后，与基准发生互换

// 一轮排序，双边循环法
func partition1(arr []int, start, end int) int {
	pivot := arr[start]
	left := start
	right := end

	for left != right {
		for left < right && pivot < arr[right] {
			right--
		}
		for left < right && pivot > arr[left] {
			left++
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[start], arr[left] = arr[left], arr[start]

	return left
}

// 单边循环法

func partition2(arr []int, startIndex, endIndex int) int {
	mark := startIndex
	pivot := arr[startIndex]
	point := startIndex + 1

	for point < endIndex {
		if arr[point] < pivot {
			mark++
			arr[mark], arr[point] = arr[point], arr[mark]
		}
		point++
	}
	arr[startIndex], arr[mark] = arr[mark], arr[startIndex]
	return mark
}

func quickSort(arr []int) []int {
	return _quickSort(arr,0,len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left < right {
		curMidIndex := partition(arr,left,right)
		_quickSort(arr,left,curMidIndex-1)
		_quickSort(arr,curMidIndex+1,right)

	}
	return arr
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := left + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr,i,index)
			index++
		}
	}
	swap(arr,pivot,index-1)
	return index-1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
