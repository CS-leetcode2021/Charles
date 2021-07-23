package __sort

import "fmt"

/**
 *  @ClassName:6_bucket
 *  @Description:桶排序
 *  @Author:jackey
 *  @Create:2021/7/23 下午7:14
 */

func main() {
	fmt.Println("桶排序学习与演示")

	arr := []int{1, 5, 5, 9, 8, 8, 6, 4, 4, 9, 3, 2, 7, 8, 5, 1, 2}
	Show(arr)

	newarr := tong(arr)

	Show(newarr)

}

//桶排序
//数组下标为排序元素的值，数组值为元素出现的个数
func tong(arr []int) []int {
	t := make([]int, 10)
	for _, val := range arr {
		t[val]++
	}
	res := make([]int, 0, len(arr))
	for index, val := range t {
		//循环把排序元素添加到新的数组中
		for ; val > 0; val-- {
			res = append(res, index)
		}
	}
	return res

}

//Show 显示数组内容
func Show(arr []int) {
	for _, item := range arr {
		fmt.Printf("%d ", item)
	}
	fmt.Println("")
}