package main

import "fmt"

/**
 *  @ClassName:bin_merge
 *  @Description:归并排序 pingcap面试题
 *  @Author:jackey
 *  @Create:2021/5/24 下午7:48
 */

// 合并 [l,r] 两部分数据，mid 左半部分的终点，mid + 1 是右半部分的起点
func merge(arr []int, l int, mid int, r int) {
	// 因为需要直接修改 arr 数据，这里首先复制 [l,r] 的数据到新的数组中，用于赋值操作
	temp := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		temp[i-l] = arr[i]
	}

	// 指向两部分起点
	// 相当于两个数组进行排序
	left := l
	right := mid + 1
	// 每次都要进行互换操作
	for i := l; i <= r; i++ {
		// 左边的点超过中点，说明只剩右边的数据
		if left > mid {
			arr[i] = temp[right-l]
			right++
			// 右边的数据超过终点，说明只剩左边的数据
		} else if right > r {
			arr[i] = temp[left-l]
			left++
			// 左边的数据大于右边的数据，选小的数字
		} else if temp[left-l] > temp[right-l] {
			arr[i] = temp[right-l]
			right++
		} else {
			// 右边的数据大于左边的数据，选择左边数据进行添加
			arr[i] = temp[left-l]
			left++
		}
	}
}

/*先分后并*/
func MergeSort(arr []int, l int, r int) {
	if l >= r {
		return
	}

	// 递归向下
	mid := (r + l) / 2
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)
	// 归并向上
	merge(arr, l, mid, r)
}

func main() {
	arr := []int{3, 1, 2, 5, 6, 43, 4}
	res := sortArray(arr)

	fmt.Println(res)
}

// 第二种实现
func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2

	left := sortArray(nums[0:mid])
	right := sortArray(nums[mid:])
	return merge2(left, right)
}

func merge2(l, r []int) []int {
	i, j := 0, 0
	m, n := len(l), len(r)
	var res []int
	for i < m && j < n {
		if l[i] < r[j] {
			res = append(res, l[i])
			i++
		} else {
			res = append(res, r[j])
			j++
		}
	}
	res = append(res,l[i:]...)
	res = append(res,r[j:]...)
	return res
}
