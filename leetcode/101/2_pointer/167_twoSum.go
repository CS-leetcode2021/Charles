package __pointer

/**
 *  @ClassName:167_twoSum
 *  @Description:easy	区间搜索
 *  @Author:jackey
 *  @Create:2021/7/22 上午10:33
 */

// 双指针，左右相反方向进行判定
// 92/100
func twoSum(numbers []int, target int) []int {
	n := len(numbers)

	for i, j := 0, n-1; i < j; {
		if numbers[i]+numbers[j] == target {
			return []int{i, j}
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	return nil
}
